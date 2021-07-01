package main

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/config"
	usersApi "golang-web-project-framework/modules/users/api"
	booksApi "golang-web-project-framework/modules/books/api"
)

func main() {
	router := gin.Default()
	usersApi.InitRouter(router)
	booksApi.InitRouter(router)
	router.Run(config.Config.Port)
}
