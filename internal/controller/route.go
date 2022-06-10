package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/tes", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Jagad!")
	})

	return e
}
