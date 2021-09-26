package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/masayukinii1011/gin-api/pkg/domain"
	"github.com/masayukinii1011/gin-api/pkg/handler"
)

var Router *gin.Engine

func InitRouter() {
	article := domain.New()
	router := gin.Default()
	router.GET("/article", handler.ArticlesGet(article))
	router.POST("/article", handler.ArticlePost(article))
	router.Run(":3000")
	Router = router
}
