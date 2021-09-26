package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/masayukinii1011/gin-api/pkg/controller"
	"github.com/masayukinii1011/gin-api/pkg/domain"
)

var Router *gin.Engine

func InitRouter() {
	//article := domain.New()
	router := gin.Default()
	//router.GET("/article", controller.GetArticles(article))
	//router.POST("/article", controller.PostArticle(article))
	router.GET("/articles", controller.GetArticles())
	router.POST("/articles", controller.PostArticle())
	router.Run(":3000")
	Router = router
}
