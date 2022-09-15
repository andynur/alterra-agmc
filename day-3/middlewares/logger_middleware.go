package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `[${time_custom}] ${status} - ${method} ${host}${uri} | ${latency_human}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	})
}
