package api

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/config"
)

func InitRouter(router *gin.Engine) {
	apiV1 := router.Group(config.Config.ApiVersion)
	usersGroup := apiV1.Group("/users")
	{
		usersGroup.POST("")
		usersGroup.DELETE("/:userId")
		usersGroup.PUT("/:userId")
		usersGroup.GET("")
		usersGroup.GET("/:userId")
	}
}
