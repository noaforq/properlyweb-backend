package gateways

import "properly.jp/properlyweb-backend/entities"

type ArticleGateway struct {
	SQLHandler
}

func (gateway *ArticleGateway) FindById(id int) (article entities.Article, err error) {
	if err = gateway.Find(&article, id).Error; err != nil {
		return
	}
	return
}

func (gateway *ArticleGateway) List() (articles entities.Articles, err error) {
	if err = gateway.Find(&articles).Error; err != nil {
		return
	}
	return
}
