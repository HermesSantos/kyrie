package main

import (
	"kyrie/controllers/quotes"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	quotesController, err := quotes.NewQuotesController("./data/quotes.json")
	if err != nil {
		panic("Error creating quotes controller: " + err.Error())
	}

	e.GET("/quotes", quotesController.GetAllQuotes)
	e.GET("/quotes/:id", quotesController.GetQuoteByID)
	e.GET("/quotes/random", quotesController.GetRandomQuote)

	e.GET("/ping", func(c *echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
