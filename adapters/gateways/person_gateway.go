package gateways

import "properly.jp/properlyweb-backend/entities"

type PersonGateway struct {
	SQLHandler
}

func (gateway *PersonGateway) FindById(id int) (person entities.Person, err error) {
	if err = gateway.Find(&person, id).Error; err != nil {
		return
	}
	return
}
