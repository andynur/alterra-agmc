package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"rest-echo/lib"
	"rest-echo/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	books, err := ReadJsonBooksData()
	if err != nil {
		return ErrorResponse(c, err, http.StatusInternalServerError)
	}

	return SuccessResponse(c, books)
}

func ShowBookController(c echo.Context) error {
	books, err := ReadJsonBooksData()
	if err != nil {
		return ErrorResponse(c, err, http.StatusInternalServerError)
	}

	bookId, _ := strconv.Atoi(c.Param("id"))
	if bookId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	var book models.BookData
	for i := 0; i < len(books); i++ {
		if books[i].Id == bookId {
			book = books[i]
		}
	}

	if book.Id == 0 {
		return ErrorResponse(c, errors.New("book not found "), http.StatusNotFound)
	}

	return SuccessResponse(c, book)
}

func StoreBookController(c echo.Context) error {
	var req lib.StoreBookRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, req)
}

func UpdateBookController(c echo.Context) error {
	books, err := ReadJsonBooksData()
	if err != nil {
		return ErrorResponse(c, err, http.StatusInternalServerError)
	}

	bookId, _ := strconv.Atoi(c.Param("id"))
	if bookId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	var book models.BookData
	for i := 0; i < len(books); i++ {
		if books[i].Id == bookId {
			book = books[i]
		}
	}

	if book.Id == 0 {
		return ErrorResponse(c, errors.New("book not found "), http.StatusNotFound)
	}

	var req lib.UpdatedBookRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	// update book by request field
	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}
	if req.Country != "" {
		book.Country = req.Country
	}
	if req.Language != "" {
		book.Language = req.Language
	}
	if req.Pages != 0 {
		book.Pages = req.Pages
	}
	if req.ImageLink != "" {
		book.ImageLink = req.ImageLink
	}
	if req.Link != "" {
		book.Link = req.Link
	}
	if req.Year != 0 {
		book.Year = req.Year
	}

	return SuccessResponse(c, book)
}

func DeleteBookController(c echo.Context) error {
	books, err := ReadJsonBooksData()
	if err != nil {
		return ErrorResponse(c, err, http.StatusInternalServerError)
	}

	bookId, _ := strconv.Atoi(c.Param("id"))
	if bookId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	var book models.BookData
	for i := 0; i < len(books); i++ {
		if books[i].Id == bookId {
			book = books[i]
		}
	}

	if book.Id == 0 {
		return ErrorResponse(c, errors.New("book not found "), http.StatusNotFound)
	}

	// reset book value
	fmt.Println("delete book #", book.Id)
	book = models.BookData{}

	return SuccessResponse(c, nil)
}
