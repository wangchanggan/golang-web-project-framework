package api

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/src/config"
	"golang-web-project-framework/src/modules/user/api/handlers"
)

var userHandlers = new(handlers.UserHandlers)

func InitRouter(router *gin.Engine) {
	apiV1 := router.Group(config.Config.ApiVersion)
	userGroup := apiV1.Group("/users")
	{
		userGroup.POST("", userHandlers.AddUser)
		userGroup.DELETE("/:userId", userHandlers.DeleteUserById)
		userGroup.PUT("/:userId", userHandlers.UpdateUserById)
		userGroup.GET("", userHandlers.GetUsers)
		userGroup.GET("/:userId", userHandlers.GetUserById)
	}
}
