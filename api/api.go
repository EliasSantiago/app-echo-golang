package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EliasFonseca/app-echo-golang/api/errors"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
)

type API struct{}

type BooksParams struct {
	Offset int `schema:"offset"`
	Limit  int `schema:"limit"`
}

type PostBook struct {
	Title string `json:"title"`
}

type UpdateBook struct {
	Title string `json:"title"`
}

var (
	books   = []string{"Book 1", "Book 2", "Book 3"}
	decoder = schema.NewDecoder()
)

func (a *API) getBooks(c echo.Context) error {
	params := &BooksParams{}
	err := decoder.Decode(params, c.QueryParams())
	if err != nil {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	if params.Offset > len(books) || params.Offset < 0 {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	if params.Limit < 0 || params.Limit > len(books) {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}
	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(books)
	}
	return c.JSON(http.StatusOK, books[from:to])
}

func (a *API) getBook(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	index := id - 1
	if index < 0 || index > len(books)-1 {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	return c.JSON(http.StatusOK, books[index])
}

func (a *API) postBook(c echo.Context) error {
	book := &PostBook{}
	err := json.NewDecoder(c.Request().Body).Decode(book)
	if err != nil {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	books = append(books, book.Title)
	return c.JSON(http.StatusOK, book.Title)
}

func (a *API) updateBook(c echo.Context) error {
	idParam := c.Param("id")
	book := &UpdateBook{}
	err := json.NewDecoder(c.Request().Body).Decode(book)
	if err != nil {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	index := id - 1
	if index < 0 || index > len(books)-1 {
		return errors.InvalidParameters(c, err, "api.books.get")
	}
	books[id-1] = book.Title
	return c.JSON(http.StatusOK, books)
}
