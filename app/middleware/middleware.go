package middleware

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"strings"
)

func RegisterMiddleware(ctx context.Context, router *mux.Router) {

	router.Use(Recovery)

	if viper.GetBool("MW_PROMETHEUS_ENABLED") {
		prometheus := NewPrometheusMiddleware(
			PrometheusMiddlewareConfig{ServiceName: viper.GetString("APP_NAME")},
		)
		router.Use(prometheus.InstrumentHandlerDuration)
	}

	if viper.GetBool("MW_CORS_ENABLED") {
		var opts []CORSOption
		opts = append(opts, MaxAge(viper.GetInt("MW_CORS_MAXAGE")))
		opts = append(opts, AllowedOrigins(strings.Split(viper.GetString("MW_CORS_ALLOWORIGINS"), ",")))
		opts = append(opts, AllowedHeaders(strings.Split(viper.GetString("MW_CORS_ALLOWHEADERS"), ",")))
		opts = append(opts, AllowedMethods(strings.Split(viper.GetString("MW_CORS_ALLOWMETHODS"), ",")))
		opts = append(opts, ExposedHeaders(strings.Split(viper.GetString("MW_CORS_EXPOSEHEADERS"), ",")))
		if viper.GetBool("MW_CORS_ALLOWCREDENTIALS") {
			opts = append(opts, AllowCredentials())
		}
		if viper.GetBool("MW_CORS_IGNOREOPTIONS") {
			opts = append(opts, IgnoreOptions())
		}

		router.Use(CORS(opts...))
	}
}
