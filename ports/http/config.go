package http

type HTTPConfig struct {
	Port         int      `default:"" envconfig:"APP_HTTP_ADDRESS"`
	Addr         string   `default:"8080" envconfig:"APP_HTTP_PORT"`
	ExcludedPath []string `default:"/metrics,/alive,/ready" envconfig:"APP_HTTP_EXCLUDED_PATH"`
}
