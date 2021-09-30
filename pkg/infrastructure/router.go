package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/masayukinii1011/gin-api/pkg/interfaces/controllers"
)

var Router *gin.Engine

func InitRouter() {
	articleController := controllers.NewArticleController()

	router := gin.Default()
	router.GET("/articles", func(c *gin.Context) { articleController.GetArticles(c) })
	router.POST("/articles", func(c *gin.Context) { articleController.PostArticle(c) })
	router.Run(":8081")
	Router = router
}
