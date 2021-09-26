package handler

import (
	"net/http"

	"github.com/masayukinii1011/gin-api/pkg/domain"

	"github.com/gin-gonic/gin"
)

func ArticlesGet(articles *domain.Articles) gin.HandlerFunc {
    return func(c *gin.Context) {
        result := articles.GetAll()
        c.JSON(http.StatusOK, result)
    }
}

type ArticlePostRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

func ArticlePost(post *domain.Articles) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := ArticlePostRequest{}
        c.Bind(&requestBody)

        item := domain.Item{
            Title:       requestBody.Title,
            Description: requestBody.Description,
        }
        post.Add(item)

        c.Status(http.StatusNoContent)
    }
}