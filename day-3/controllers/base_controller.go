package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, DefaultResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, err error, httpCode int) error {
	fmt.Println("ERROR : ", err)
	return c.JSON(httpCode, DefaultResponse{
		Code:    httpCode,
		Message: err.Error(),
		Data:    nil,
	})
}
