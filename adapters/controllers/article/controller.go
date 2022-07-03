package article

import (
	"strconv"

	"github.com/labstack/echo/v4"
	gateways "properly.jp/properlyweb-backend/adapters/gateways"
	usecases "properly.jp/properlyweb-backend/usecases/interactors/article"
)

type ArticleController struct {
	Interactor usecases.ArticleInteractor
}

func InitArticleController(sqlHandler gateways.SQLHandler) *ArticleController {
	return &ArticleController{
		Interactor: usecases.ArticleInteractor{
			ArticlePort: &gateways.ArticleGateway{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *ArticleController) FindById(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := controller.Interactor.FindById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, article)
	return
}

func (controller *ArticleController) List(c echo.Context) (err error) {
	articles, err := controller.Interactor.List()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, articles)
	return
}
