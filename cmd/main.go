package main

import (
	"github.com/masayukinii1011/gin-api/pkg/handler"

	"github.com/masayukinii1011/gin-api/pkg/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	article := article.New()
	r := gin.Default()
	r.GET("/article", handler.ArticlesGet(domain.article.Articles))
	r.POST("/article", handler.ArticlePost(domain.article.Articles))
	r.Run(":3000")
}