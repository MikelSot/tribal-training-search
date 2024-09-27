package bootstrap

import "os"

const _nameAppDefault = "ms-search"

const _portDefault = ":3000"

const (
	_allowOriginsDefault = "*"
	_allowMethodsDefault = "GET,POST"
)

func getApplicationName() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		return _nameAppDefault
	}

	return appName
}

func getPort() string {
	port := os.Getenv("FIBER_PORT")
	if port == "" {
		return _portDefault
	}

	return port
}

func getAllowOrigins() string {
	allowedOrigins := os.Getenv("ALLOW_ORIGINS")
	if allowedOrigins == "" {
		return _allowOriginsDefault
	}

	return allowedOrigins
}

func getAllowMethods() string {
	allowedMethods := os.Getenv("ALLOW_METHODS")
	if allowedMethods == "" {
		return _allowMethodsDefault
	}

	return allowedMethods
}

func getItunesRoute() string {
	routeStatistic := os.Getenv("ITUNES_ROUTE")
	if routeStatistic == "" {
		return ""
	}

	return routeStatistic
}

func getChartsLyricsRoute() string {
	routeStatistic := os.Getenv("CHARTS_LYRICS_ROUTE")
	if routeStatistic == "" {
		return ""
	}

	return routeStatistic
}
