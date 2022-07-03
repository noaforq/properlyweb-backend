package person

import (
	"strconv"

	"github.com/labstack/echo/v4"
	gateways "properly.jp/properlyweb-backend/adapters/gateways"
	usecases "properly.jp/properlyweb-backend/usecases/interactors/person"
)

type PersonController struct {
	Interactor usecases.PersonInteractor
}

func InitPersonController(sqlHandler gateways.SQLHandler) *PersonController {
	return &PersonController{
		Interactor: usecases.PersonInteractor{
			PersonPort: &gateways.PersonGateway{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *PersonController) FindById(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	person, err := controller.Interactor.FindById(id)
	if person.Id == 0 {
		c.JSON(404, "Not Found")
		return
	}
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, person)
	return
}
