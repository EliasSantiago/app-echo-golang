package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	public := e.Group("")

	public.GET("/books", a.getBooks)
	public.GET("/books/:id", a.getBook)
	public.POST("/books", a.postBook)
	public.PUT("/books/:id", a.updateBook)
}
