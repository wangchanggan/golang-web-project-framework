package main

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/src/config"
	bookApi "golang-web-project-framework/src/modules/book/api"
	userApi "golang-web-project-framework/src/modules/user/api"
)

func main() {
	router := gin.Default()
	userApi.InitRouter(router)
	bookApi.InitRouter(router)
	router.Run(config.Config.Port)
}
