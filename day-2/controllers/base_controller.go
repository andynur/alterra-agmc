package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-echo/lib"
	"rest-echo/models"

	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, lib.DefaultResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, err error, httpCode int) error {
	fmt.Println("ERROR : ", err)
	return c.JSON(httpCode, lib.DefaultResponse{
		Code:    httpCode,
		Message: err.Error(),
		Data:    nil,
	})
}

func ReadJsonBooksData() (data []models.BookData, err error) {
	file, err := ioutil.ReadFile("storage/books.json")
	if err != nil {
		return data, err
	}

	books := models.Books{}
	err = json.Unmarshal([]byte(file), &books)
	if err != nil {
		return data, err
	}

	return books.Data, nil
}
