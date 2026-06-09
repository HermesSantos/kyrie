package router

import "github.com/labstack/echo/v5"

func InitRoutes(e *echo.Group) {
	InitEjaculatoriesRoute(e)
	InitQuotesRoute(e)
}
