package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SkipperData struct {
	Method  string
	UrlPath string
}

func JwtMiddleware() echo.MiddlewareFunc {
	secretKey := os.Getenv("JWT_SECRET")

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(secretKey),
		Skipper: func(c echo.Context) bool {
			for _, item := range whiteListRoutes() {
				if c.Request().URL.Path == item.UrlPath && c.Request().Method == item.Method {
					return true
				}
			}

			return false
		},
	})
}

func whiteListRoutes() []SkipperData {
	return []SkipperData{
		{"POST", "/v1/auth/login"},
		{"GET", "/v1/books"},
		{"GET", "/v1/books/:id"},
		{"POST", "/v1/users"},
	}
}
