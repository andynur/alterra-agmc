package routes

import (
	"rest-echo/controllers"
	"rest-echo/middlewares"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {
	// group route to v1
	v1 := e.Group("v1")

	// setup middlewares
	v1.Use(middlewares.LoggerMiddleware())
	v1.Use(middlewares.JwtMiddleware())

	// auth routes
	v1.POST("/auth/login", controllers.AuthLoginController)

	// books routes
	v1.GET("/books", controllers.GetBooksController)
	v1.GET("/books/:id", controllers.ShowBookController)
	v1.POST("/books", controllers.StoreBookController)
	v1.PUT("/books/:id", controllers.UpdateBookController)
	v1.DELETE("/books/:id", controllers.DeleteBookController)

	// users routes
	v1.GET("/users", controllers.GetUsersController)
	v1.GET("/users/:id", controllers.ShowUserController)
	v1.POST("/users", controllers.StoreUserController)
	v1.PUT("/users/:id", controllers.UpdateUserController)
	v1.DELETE("/users/:id", controllers.DeleteUserController)
}
