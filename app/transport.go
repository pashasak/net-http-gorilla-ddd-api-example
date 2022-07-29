package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"

	"net-http-gorilla-ddd-example/utils"
	"net/http"
	"net/http/pprof"
)

func NewRouter(ctx context.Context) *mux.Router {

	versionHandler := func(w http.ResponseWriter, r *http.Request) {
		buildInfo := BuildInfo{
			BuildTime:    BuildTime,
			BuildBranch:  BuildBranch,
			BuildSummary: BuildSummary,
			BuildCommit:  BuildCommit,
		}
		utils.JsonPresenter(w, buildInfo)
	}

	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	router := mux.NewRouter()

	// System Routes
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/health", healthHandler).Methods("GET")
	router.HandleFunc("/version", versionHandler).Methods("GET")
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		//httpSwagger.URL("docs/swagger.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	return router
}
