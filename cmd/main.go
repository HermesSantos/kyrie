package main

import (
	"kyrie/controllers/ejaculatories"
	"kyrie/controllers/quotes"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	en := echo.New()
	en.Use(middleware.RequestLogger())
	en.Use(middleware.ContextTimeout(60 * time.Second))
	en.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20.0)))
	en.Pre(middleware.RemoveTrailingSlash())

	e := en.Group("/api")

	// quotes
	quotesController, err := quotes.NewQuotesController("./data/quotes.json")
	if err != nil {
		panic("Error creating quotes controller: " + err.Error())
	}

	e.GET("/quotes", quotesController.GetAllQuotes)
	e.GET("/quotes/:id", quotesController.GetQuoteByID)
	e.GET("/quotes/random", quotesController.GetRandomQuote)

	// ejaculatories
	ejaculatoriesController, err := ejaculatories.NewEjaculatoriesController("./data/ejaculatories.json")
	if err != nil {
		panic("Error creating ejaculatories controller: " + err.Error())
	}

	e.GET("/ejaculatories", ejaculatoriesController.GetAllEjaculatories)
	e.GET("/ejaculatories/random", ejaculatoriesController.GetRandomEjaculatory)
	e.GET("/ejaculatories/categories", ejaculatoriesController.GetEjaculatoriesCategories)
	e.GET("/ejaculatories/category/:category", ejaculatoriesController.GetEjaculatoriesByCategory)
	e.GET("/ejaculatories/:id", ejaculatoriesController.GetEjaculatoryByID)

	// test
	e.GET("/ping", func(c *echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	if err := en.Start(":1323"); err != nil {
		en.Logger.Error("failed to start server", "error", err)
	}
}
