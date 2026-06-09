package quotes

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuotesController struct {
	quotes []Quote
}

func NewQuotesController(path string) (*QuotesController, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var quotes []Quote

	err = json.Unmarshal(file, &quotes)
	if err != nil {
		return nil, err
	}

	return &QuotesController{
		quotes: quotes,
	}, nil
}

func (qc *QuotesController) GetAllQuotes(c *echo.Context) error {
	return c.JSON(200, qc.quotes)
}

func (qc *QuotesController) GetQuoteByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(404, "Not found")
	}

	for _, quote := range qc.quotes {
		if quote.ID == id {
			return c.JSON(200, quote)
		}
	}

	return c.JSON(404, fmt.Sprintf("Quote with id %d not found", id))
}

func (qc *QuotesController) GetRandomQuote(c *echo.Context) error {
	randomIndex := rand.IntN(len(qc.quotes))

	return c.JSON(200, qc.quotes[randomIndex])
}
