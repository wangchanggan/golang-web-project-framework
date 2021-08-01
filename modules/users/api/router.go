package api

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/config"
	"golang-web-project-framework/modules/users/api/handlers"
)

var usersHandlers = new(handlers.UsersHandlers)

func InitRouter(router *gin.Engine) {
	apiV1 := router.Group(config.Config.ApiVersion)
	usersGroup := apiV1.Group("/users")
	{
		usersGroup.POST("", usersHandlers.AddUser)
		usersGroup.DELETE("/:userId")
		usersGroup.PUT("/:userId")
		usersGroup.GET("")
		usersGroup.GET("/:userId")
	}
}
