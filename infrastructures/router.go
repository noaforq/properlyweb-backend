package infrastructures

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	routes "properly.jp/properlyweb-backend/infrastructures/routes"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	routes.InitArticle(e)

	e.Logger.Fatal(e.Start(":1323"))
}
