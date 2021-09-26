package controller

import (
	"net/http"
	"github.com/masayukinii1011/gin-api/pkg/domain"
    "github.com/masayukinii1011/gin-api/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
    ArticleRepository usecase.ArticleRepository
}

type ArticlePostRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

func GetArticles(controller ArticleController) gin.HandlerFunc {
    return func(c *gin.Context) {
        result := controller.ArticleRepository.GetAll()
        c.JSON(http.StatusOK, result)
    }
}

func PostArticle(controller ArticleController) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := ArticlePostRequest{}
        c.Bind(&requestBody)
        item := domain.Article{
            Title:       requestBody.Title,
            Description: requestBody.Description,
        }
        controller.ArticleRepository.Add(item)
        c.Status(http.StatusNoContent)
    }
}