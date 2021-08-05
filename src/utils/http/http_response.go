package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang-web-project-framework/src/utils"
)

func Response(c *gin.Context, httpStatusCode int, obj interface{}) {
	c.JSON(httpStatusCode, obj)
}

func ResponseWithOk(c *gin.Context, content interface{}) {
	Response(c, http.StatusOK, gin.H{
		utils.ErrorCode: utils.ErrorCodeSuccess,
		utils.ErrorDesc: utils.ErrorDescSuccess,
		utils.Content: content})
}

func ResponseWithFailed(c *gin.Context, utilsErr *utils.Error)  {
	Response(c, http.StatusOK, utilsErr)
}

func ResponseWithParameterError(c *gin.Context, errorDesc string)  {
	Response(c, http.StatusBadRequest, utils.NewError(utils.ErrorCodeParameterError, errorDesc))
}
