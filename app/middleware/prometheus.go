package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	dflBuckets = []float64{0.001, 0.01, 0.05, 0.1, 0.25, 0.5, 0.75, 1}
)

const (
	requestName = "requests_total"
	latencyName = "request_duration_milliseconds"
	errsName    = "errors_total"
)

// Opts specifies options how to create new PrometheusMiddleware.
type PrometheusMiddlewareConfig struct {
	// Buckets specifies an custom buckets to be used in request histograpm.
	Buckets     []float64
	ServiceName string
}

// PrometheusMiddleware specifies the metrics that is going to be generated
type PrometheusMiddleware struct {
	request *prometheus.CounterVec
	latency *prometheus.HistogramVec
	errs    *prometheus.CounterVec
}

// NewPrometheusMiddleware creates a new PrometheusMiddleware instance
func NewPrometheusMiddleware(opts PrometheusMiddlewareConfig) *PrometheusMiddleware {
	var prometheusMiddleware PrometheusMiddleware

	prometheusMiddleware.request = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        requestName,
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: prometheus.Labels{"service": opts.ServiceName},
		},
		[]string{"method", "status", "path"},
	)

	if err := prometheus.Register(prometheusMiddleware.request); err != nil {
		log.Println("prometheusMiddleware.request was not registered:", err)
	}

	buckets := opts.Buckets
	if len(buckets) == 0 {
		buckets = dflBuckets
	}

	prometheusMiddleware.latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        latencyName,
		Help:        "How long it took to process the request, partitioned by status code, method and HTTP path.",
		ConstLabels: prometheus.Labels{"service": opts.ServiceName},
		Buckets:     buckets,
	},
		[]string{"code", "method", "path"},
	)

	if err := prometheus.Register(prometheusMiddleware.latency); err != nil {
		log.Println("prometheusMiddleware.latency was not registered:", err)
	}

	prometheusMiddleware.errs = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        errsName,
		Help:        "How mane errors occurred",
		ConstLabels: prometheus.Labels{"service": opts.ServiceName},
	},
		[]string{},
	)

	return &prometheusMiddleware
}

// InstrumentHandlerDuration is a middleware that wraps the http.Handler and it record
// how long the handler took to run, which path was called, and the status code.
// This method is going to be used with gorilla/mux.
func (p *PrometheusMiddleware) InstrumentHandlerDuration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()

		delegate := &responseWriterDelegator{ResponseWriter: w}
		rw := delegate

		next.ServeHTTP(rw, r) // call original

		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		code := sanitizeCode(delegate.status)
		method := sanitizeMethod(r.Method)

		go p.request.WithLabelValues(
			code,
			method,
			path,
		).Inc()

		go p.latency.WithLabelValues(
			code,
			method,
			path,
		).Observe(float64(time.Since(begin)) / float64(time.Second))

	})
}

type responseWriterDelegator struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (r *responseWriterDelegator) WriteHeader(code int) {
	r.status = code
	r.wroteHeader = true
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseWriterDelegator) Write(b []byte) (int, error) {
	if !r.wroteHeader {
		r.WriteHeader(http.StatusOK)
	}
	n, err := r.ResponseWriter.Write(b)
	r.written += int64(n)
	return n, err
}

func sanitizeMethod(m string) string {
	return strings.ToLower(m)
}

func sanitizeCode(s int) string {
	return strconv.Itoa(s)
}
