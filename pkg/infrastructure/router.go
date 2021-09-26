package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/masayukinii1011/gin-api/pkg/interfaces/controllers"
)

var Router *gin.Engine

func InitRouter() {
	router := gin.Default()
	router.GET("/articles", func(c *gin.Context) { controllers.GetArticles(c) })
	router.POST("/articles", func(c *gin.Context) { controllers.PostArticle(c) })
	router.Run(":3000")
	Router = router
}
