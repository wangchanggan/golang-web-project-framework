package main

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/src/config"
	usersApi "golang-web-project-framework/src/modules/users/api"
	booksApi "golang-web-project-framework/src/modules/books/api"
)

func main() {
	router := gin.Default()
	usersApi.InitRouter(router)
	booksApi.InitRouter(router)
	router.Run(config.Config.Port)
}
