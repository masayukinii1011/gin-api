package main

import (
	"github.com/masayukinii1011/gin-api/handler"

	"github.com/masayukinii1011/gin-api/article"

	"github.com/gin-gonic/gin"
)

func main() {
    article := article.New()
    r := gin.Default()
    r.GET("/article", handler.ArticlesGet(article))
    r.POST("/article", handler.ArticlePost(article))

    r.Run() // listen and serve on 0.0.0.0:8080

}