package router

import (
	"kyrie/controllers/quotes"

	"github.com/labstack/echo/v5"
)

func InitQuotesRoute(e *echo.Group) {

	quotesController, err := quotes.NewQuotesController("./data/quotes.json")
	if err != nil {
		panic("Error creating quotes controller: " + err.Error())
	}

	e.GET("/quotes", quotesController.GetAllQuotes)
	e.GET("/quotes/:id", quotesController.GetQuoteByID)
	e.GET("/quotes/random", quotesController.GetRandomQuote)
}
