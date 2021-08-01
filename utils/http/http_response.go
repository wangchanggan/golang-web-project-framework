package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	utilsError "golang-web-project-framework/utils/error"
)

func Response(c *gin.Context, httpStatusCode int, obj interface{}) {
	c.JSON(httpStatusCode, obj)
}

func ResponseWithOk(c *gin.Context, content interface{}) {
	Response(c, http.StatusOK, gin.H{
		utilsError.ErrorCode: utilsError.ErrorCodeSuccess,
		utilsError.ErrorDesc: utilsError.ErrorDescSuccess,
		utilsError.Content: content})
}

func ResponseWithFailed(c *gin.Context, utilsErr *utilsError.Error)  {
	Response(c, http.StatusOK, utilsErr)
}

func ResponseWithBadRequest(c *gin.Context, errorDesc string)  {
	Response(c, http.StatusBadRequest, utilsError.NewError(utilsError.ErrorCodeParameterError, errorDesc))
}
