package main

import (
	"os"
	"rest-echo/config"
	"rest-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// init echo framework
	e := echo.New()

	// setup validator
	e.Validator = config.InitValidator()

	// init gorm database
	config.InitDB()

	// setup routes
	routes.ApiRoutes(e)

	// run server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
