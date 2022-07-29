package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Init() {
	// Set default configurations
	setDefaults()

	// Select the .env file
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}

	// Automatically refresh environment variables
	viper.AutomaticEnv()
}

func GetDBConnectString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_DATABASE"))

	//return fmt.Sprintf("host=%s dbname=%s port=%d user=%s password=%s sslmode=disable application_name=%s",
	//	viper.GetString("DB_HOST"),
	//	viper.GetString("DB_DATABASE"),
	//	viper.GetInt("DB_PORT"),
	//	viper.GetString("DB_USERNAME"),
	//	viper.GetString("DB_PASSWORD"),
	//	viper.GetString("APP_NAME"))
}

func setDefaults() {
	// Set default App configuration
	viper.SetDefault("APP_ADDR", ":8069")
	viper.SetDefault("APP_MAXHEADERBYTES", 16777216) // 16 * 1024 * 1024
	viper.SetDefault("APP_ENV", "local")

	// Set default database configuration
	viper.SetDefault("DB_DRIVER", "pgx")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USERNAME", "admin")
	viper.SetDefault("DB_PASSWORD", "masterkey")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_DATABASE", "db")

	// Set default CORS middleware configuration
	viper.SetDefault("MW_CORS_ENABLED", true)
	viper.SetDefault("MW_CORS_ALLOWORIGINS", "*")
	viper.SetDefault("MW_CORS_ALLOWMETHODS", "GET,POST,HEAD,PUT,DELETE,PATCH")
	viper.SetDefault("MW_CORS_ALLOWHEADERS", "X-Requested-With, X-Actual-SupplierId, X-SupplierId, X-User-Id, X-Debug-Mode, X-Debug-Supplier-Id, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header, x-office-api")
	viper.SetDefault("MW_CORS_ALLOWCREDENTIALS", false)
	viper.SetDefault("MW_CORS_IGNOREOPTIONS", false)
	viper.SetDefault("MW_CORS_EXPOSEHEADERS", "")
	viper.SetDefault("MW_CORS_MAXAGE", 0)

	viper.SetDefault("MW_PROMETHEUS_ENABLED", true)

	//// Set default Custom Access Logger middleware configuration
	//viper.SetDefault("MW_ACCESS_LOGGER_ENABLED", false)
	//viper.SetDefault("MW_ACCESS_LOGGER_TYPE", "console")
	//viper.SetDefault("MW_ACCESS_LOGGER_FILENAME", "access.log")
	//viper.SetDefault("MW_ACCESS_LOGGER_MAXSIZE", 500)
	//viper.SetDefault("MW_ACCESS_LOGGER_MAXAGE", 28)
	//viper.SetDefault("MW_ACCESS_LOGGER_MAXBACKUPS", 3)
	//viper.SetDefault("MW_ACCESS_LOGGER_LOCALTIME", false)
	//viper.SetDefault("MW_ACCESS_LOGGER_COMPRESS", false)
	//
	//// Set default Force HTTPS middleware configuration
	//viper.SetDefault("MW_FORCE_HTTPS_ENABLED", false)
	//
	//// Set default Force trailing slash middleware configuration
	//viper.SetDefault("MW_FORCE_TRAILING_SLASH_ENABLED", false)
	//
	//// Set default HSTS middleware configuration
	//viper.SetDefault("MW_HSTS_ENABLED", false)
	//viper.SetDefault("MW_HSTS_MAXAGE", 31536000)
	//viper.SetDefault("MW_HSTS_INCLUDESUBDOMAINS", true)
	//viper.SetDefault("MW_HSTS_PRELOAD", false)
	//
	//// Set default Suppress WWW middleware configuration
	//viper.SetDefault("MW_SUPPRESS_WWW_ENABLED", true)
	//
	//// Set default Fiber Cache middleware configuration
	//viper.SetDefault("MW_FIBER_CACHE_ENABLED", false)
	//viper.SetDefault("MW_FIBER_CACHE_EXPIRATION", "1m")
	//viper.SetDefault("MW_FIBER_CACHE_CACHECONTROL", false)
	//
	//// Set default Fiber Compress middleware configuration
	//viper.SetDefault("MW_FIBER_COMPRESS_ENABLED", false)
	//viper.SetDefault("MW_FIBER_COMPRESS_LEVEL", 0)
	//
	//// Set default Fiber CORS middleware configuration
	//viper.SetDefault("MW_FIBER_CORS_ENABLED", false)
	//viper.SetDefault("MW_FIBER_CORS_ALLOWORIGINS", "*")
	//viper.SetDefault("MW_FIBER_CORS_ALLOWMETHODS", "GET,POST,HEAD,PUT,DELETE,PATCH")
	//viper.SetDefault("MW_FIBER_CORS_ALLOWHEADERS", "")
	//viper.SetDefault("MW_FIBER_CORS_ALLOWCREDENTIALS", false)
	//viper.SetDefault("MW_FIBER_CORS_EXPOSEHEADERS", "")
	//viper.SetDefault("MW_FIBER_CORS_MAXAGE", 0)
	//
	//// Set default Fiber CSRF middleware configuration
	//viper.SetDefault("MW_FIBER_CSRF_ENABLED", false)
	//viper.SetDefault("MW_FIBER_CSRF_TOKENLOOKUP", "header:X-CSRF-Token")
	//viper.SetDefault("MW_FIBER_CSRF_COOKIE_NAME", "_csrf")
	//viper.SetDefault("MW_FIBER_CSRF_COOKIE_SAMESITE", "Strict")
	//viper.SetDefault("MW_FIBER_CSRF_COOKIE_EXPIRES", "24h")
	//viper.SetDefault("MW_FIBER_CSRF_CONTEXTKEY", "csrf")
	//
	//// Set default Fiber ETag middleware configuration
	//viper.SetDefault("MW_FIBER_ETAG_ENABLED", false)
	//viper.SetDefault("MW_FIBER_ETAG_WEAK", false)
	//
	//// Set default Fiber Expvar middleware configuration
	//viper.SetDefault("MW_FIBER_EXPVAR_ENABLED", false)
	//
	//// Set default Fiber Favicon middleware configuration
	//viper.SetDefault("MW_FIBER_FAVICON_ENABLED", false)
	//viper.SetDefault("MW_FIBER_FAVICON_FILE", "")
	//viper.SetDefault("MW_FIBER_FAVICON_CACHECONTROL", "public, max-age=31536000")
	//
	//// Set default Fiber Limiter middleware configuration
	//viper.SetDefault("MW_FIBER_LIMITER_ENABLED", true)
	//viper.SetDefault("MW_FIBER_LIMITER_MAX", 5)
	//viper.SetDefault("MW_FIBER_LIMITER_EXPIRATION", "1m")
	//
	//// Set default Fiber Monitor middleware configuration
	//viper.SetDefault("MW_FIBER_MONITOR_ENABLED", false)
	//
	//// Set default Fiber Pprof middleware configuration
	//viper.SetDefault("MW_FIBER_PPROF_ENABLED", false)
	//
	//// Set default Fiber Recover middleware configuration
	//viper.SetDefault("MW_FIBER_RECOVER_ENABLED", true)
	//
	//// Set default Fiber RequestID middleware configuration
	//viper.SetDefault("MW_FIBER_REQUESTID_ENABLED", false)
	//viper.SetDefault("MW_FIBER_REQUESTID_HEADER", "X-Request-ID")
	//viper.SetDefault("MW_FIBER_REQUESTID_CONTEXTKEY", "requestid")
	//
	//// Set default Fiber Logger middleware configuration
	//viper.SetDefault("MW_FIBER_LOGGER_ENABLED", true)
	//viper.SetDefault("MW_FIBER_LOGGER_FORMAT", "${pid} ${locals:requestid} ${status} - ${method} ${path}\n")
	//viper.SetDefault("MW_FIBER_LOGGER_TIMEFORMAT", "15:04:05")
	//viper.SetDefault("MW_FIBER_LOGGER_TIMEINTERVAL", 500*time.Millisecond)
	//viper.SetDefault("MW_FIBER_LOGGER_TIMEZONE", "Europe/Moscow")
	//
	//// Set  Fiber Helmet middleware configuration
	//viper.SetDefault("MW_FIBER_HELMET_ENABLED", false)
	//viper.SetDefault("MW_FIBER_HELMET_XSS_PROTECTION", "1; mode=block")
	//viper.SetDefault("MW_FIBER_HELMET_CONTENT_TYPE_NOSNIFF", "nosniff")
	//viper.SetDefault("MW_FIBER_HELMET_X_FRAMEOPTIONS", "SAMEORIGIN")
	//viper.SetDefault("MW_FIBER_HELMET_HSTS_MAX_AGE", 0)
	//viper.SetDefault("MW_FIBER_HELMET_HSTS_EXCLUDE_SUBDOMAINS", false)
	//viper.SetDefault("MW_FIBER_HELMET_CONTENT_SECURITY_POLICY", "")
	//viper.SetDefault("MW_FIBER_HELMET_CSP_REPORT_ONLY", false)
	//viper.SetDefault("MW_FIBER_HELMET_HSTS_PRELOAD_ENABLED", false)
	//viper.SetDefault("MW_FIBER_HELMET_REFERRER_POLICY", "")
	//viper.SetDefault("MW_FIBER_HELMET_PERMISSION_POLICY", "")
	//
	//// Set Fiber Prometheus middleware configurations
	//viper.SetDefault("MW_FIBER_PROMETHEUS_ENABLED", false)
	//viper.SetDefault("MW_FIBER_PROMETHEUS_SERVICE_NAME", "my-service")
}
