package static

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"rest-echo/lib/dto"
	"rest-echo/models"
)

func GetJsonBooks() (data []models.BookData, err error) {
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

func GetJsonBookById(id int) (book models.BookData, err error) {
	books, err := GetJsonBooks()
	if err != nil {
		return book, err
	}

	for i := 0; i < len(books); i++ {
		if books[i].Id == id {
			book = books[i]
		}
	}

	if book.Id == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}

func UpdateJsonBookById(id int, req dto.UpdatedBookRequest) (book models.BookData, err error) {
	book, err = GetJsonBookById(id)
	if err != nil {
		return book, err
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

	return book, nil
}

func DeleteJsonBookById(id int) (book models.BookData, err error) {
	book, err = GetJsonBookById(id)
	if err != nil {
		return book, err
	}

	// reset book value
	fmt.Println("delete book #", book.Id)
	book = models.BookData{}

	return book, nil
}
