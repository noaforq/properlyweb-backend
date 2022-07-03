package article

import "properly.jp/properlyweb-backend/entities"

type ArticlePort interface {
	FindById(id int) (entities.Article, error)
	List() (entities.Articles, error)
}

type ArticleInteractor struct {
	ArticlePort ArticlePort
}

func (interactor *ArticleInteractor) FindById(id int) (article entities.Article, err error) {
	article, err = interactor.ArticlePort.FindById(id)
	return
}

func (interactor *ArticleInteractor) List() (articles entities.Articles, err error) {
	articles, err = interactor.ArticlePort.List()
	return
}
