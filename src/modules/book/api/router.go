package api

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/src/config"
)

func InitRouter(router *gin.Engine) *gin.Engine {
	apiV1 := router.Group(config.Config.ApiVersion)
	bookGroup := apiV1.Group("/books")
	{
		bookGroup.POST("")
		bookGroup.DELETE("/:bookId")
		bookGroup.PUT("/:bookId")
		bookGroup.GET("")
		bookGroup.GET("/:bookId")
	}

	return router
}
