package controllers

import (
	"errors"
	"net/http"
	"rest-echo/lib/dto"
	"rest-echo/lib/static"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := static.GetJsonBooks()
	if err != nil {
		return ErrorResponse(c, err, http.StatusInternalServerError)
	}

	return SuccessResponse(c, books)
}

func ShowBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))
	if bookId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	book, err := static.GetJsonBookById(bookId)
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, book)
}

func StoreBookController(c echo.Context) error {
	var req dto.StoreBookRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, req)
}

func UpdateBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))
	if bookId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	var req dto.UpdatedBookRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	book, err := static.UpdateJsonBookById(bookId, req)
	if err != nil {
		return ErrorResponse(c, err, http.StatusInternalServerError)
	}

	return SuccessResponse(c, book)
}

func DeleteBookController(c echo.Context) error {
	bookId, _ := strconv.Atoi(c.Param("id"))
	if bookId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	_, err := static.DeleteJsonBookById(bookId)
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, nil)
}
