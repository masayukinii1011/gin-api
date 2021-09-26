package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/masayukinii1011/gin-api/pkg/domain"
	"github.com/masayukinii1011/gin-api/pkg/usecase"
	"net/http"
)

type ArticleController struct {
	ArticleRepository usecase.ArticleRepository
}

type ArticlePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		ArticleRepository: usecase.ArticleRepository{},
	}
}

func (controller ArticleController) GetArticles(c *gin.Context) {
	result := controller.ArticleRepository.GetAll()
	c.JSON(http.StatusOK, result)
}

func (controller ArticleController) PostArticle(c *gin.Context) {
	requestBody := ArticlePostRequest{}
	c.Bind(&requestBody)
	item := domain.Article{
		Title:       requestBody.Title,
		Description: requestBody.Description,
	}
	controller.ArticleRepository.Add(item)
	c.Status(http.StatusNoContent)
}
