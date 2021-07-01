package api

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/config"
)

func InitRouter(router *gin.Engine) *gin.Engine {
	apiV1 := router.Group(config.Config.ApiVersion)
	booksGroup := apiV1.Group("/books")
	{
		booksGroup.POST("")
		booksGroup.DELETE("/:bookId")
		booksGroup.PUT("/:bookId")
		booksGroup.GET("")
		booksGroup.GET("/:bookId")
	}

	return router
}
