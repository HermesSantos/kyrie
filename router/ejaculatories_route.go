package router

import (
	"kyrie/controllers/ejaculatories"

	"github.com/labstack/echo/v5"
)

func InitEjaculatoriesRoute(e *echo.Group) {
	ejaculatoriesController, err := ejaculatories.NewEjaculatoriesController("./data/ejaculatories.json")
	if err != nil {
		panic("Error creating ejaculatories controller: " + err.Error())
	}

	e.GET("/ejaculatories", ejaculatoriesController.GetAllEjaculatories)
	e.GET("/ejaculatories/random", ejaculatoriesController.GetRandomEjaculatory)
	e.GET("/ejaculatories/categories", ejaculatoriesController.GetEjaculatoriesCategories)
	e.GET("/ejaculatories/category/:category", ejaculatoriesController.GetEjaculatoriesByCategory)
	e.GET("/ejaculatories/:id", ejaculatoriesController.GetEjaculatoryByID)
}
