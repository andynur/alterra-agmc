package main

import (
	"os"
	"rest-echo/config"
	"rest-echo/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// init echo framework
	e := echo.New()

	// setup validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// init gorm database
	config.InitDB()

	// setup routing
	routes.SetupRoutes(e)

	// run server
	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
