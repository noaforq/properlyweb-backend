package routes

import (
	"github.com/labstack/echo/v4"
	controller "properly.jp/properlyweb-backend/adapters/controllers/person"
	"properly.jp/properlyweb-backend/infrastructures/datastores"
)

func InitPerson(e *echo.Echo) {
	controller := controller.InitPersonController(datastores.InitSQLHandler())

	e.GET("/persons/:id", func(c echo.Context) error { return controller.FindById(c) })
}
