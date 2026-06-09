package ejaculatories

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Ejaculatory struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Quote    string `json:"quote"`
}

type EjaculatoriesController struct {
	ejaculatories []Ejaculatory
}

func NewEjaculatoriesController(path string) (*EjaculatoriesController, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var ejaculatories []Ejaculatory

	err = json.Unmarshal(file, &ejaculatories)
	if err != nil {
		return nil, err
	}

	return &EjaculatoriesController{
		ejaculatories: ejaculatories,
	}, nil
}

func (ec *EjaculatoriesController) GetAllEjaculatories(c *echo.Context) error {
	return c.JSON(200, ec.ejaculatories)
}

func (ec *EjaculatoriesController) GetEjaculatoryByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(404, "Not found")
	}

	for _, v := range ec.ejaculatories {
		if v.ID == id {
			return c.JSON(200, v)
		}
	}

	return c.JSON(404, fmt.Sprintf("Ejaculatories with id %d not found", id))
}

func (ec *EjaculatoriesController) GetRandomEjaculatory(c *echo.Context) error {
	randomIndex := rand.IntN(len(ec.ejaculatories))
	return c.JSON(200, ec.ejaculatories[randomIndex])
}

func (ec *EjaculatoriesController) GetEjaculatoriesCategories(c *echo.Context) error {
	unique := make(map[string]bool)

	for _, prayer := range ec.ejaculatories {
		unique[prayer.Category] = true
	}

	categories := make([]string, 0, len(unique))

	for category := range unique {
		categories = append(categories, category)
	}

	return c.JSON(200, categories)

}

func (ec *EjaculatoriesController) GetEjaculatoriesByCategory(c *echo.Context) error {
	category := c.Param("category")

	var filtered []Ejaculatory

	for _, prayer := range ec.ejaculatories {
		if prayer.Category == category {
			filtered = append(filtered, prayer)
		}
	}

	return c.JSON(200, filtered)
}
