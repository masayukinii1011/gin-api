package usecase

import (
	"github.com/masayukinii1011/gin-api/pkg/domain"
)

type ArticleRepository struct {
	Articles domain.Articles
}
/*
func New() *domain.Articles {
	return &domain.Articles{}
}
*/
func (repository ArticleRepository) Add(newArticle domain.Article) {
	repository.Articles = append(repository.Articles, newArticle)
}

func (repository ArticleRepository) GetAll() domain.Articles {
	return repository.Articles
}
