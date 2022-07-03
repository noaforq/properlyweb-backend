package routes

import (
	"github.com/labstack/echo/v4"
	controller "properly.jp/properlyweb-backend/adapters/controllers/article"
	"properly.jp/properlyweb-backend/infrastructures/datastores"
)

func InitArticle(e *echo.Echo) {
	controller := controller.InitArticleController(datastores.InitSQLHandler())

	e.GET("/articles/:id", func(c echo.Context) error { return controller.FindById(c) })
	e.GET("/articles", func(c echo.Context) error { return controller.List(c) })
}
