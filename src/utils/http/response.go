package http

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/src/utils"
	"net/http"
)

type PageInfo struct {
	TotalNum  int64 `json:"total_num"`
	TotalSize int64 `json:"total_size"`
}

func GeneratePageInfo(pageSize, totalSize int64) PageInfo {
	var totalNum int64
	if totalSize%pageSize == 0 {
		totalNum = totalSize / pageSize
	} else {
		totalNum = (totalSize / pageSize) + 1
	}

	return PageInfo{
		TotalNum:  totalNum,
		TotalSize: totalSize,
	}
}

func Response(c *gin.Context, httpStatusCode int, obj interface{}) {
	c.JSON(httpStatusCode, obj)
}

func ResponseWithOk(c *gin.Context, content interface{}) {
	if content == nil {
		Response(c, http.StatusOK, gin.H{
			utils.ErrorCode: utils.ErrorCodeSuccess,
			utils.ErrorDesc: utils.ErrorDescSuccess})
	} else {
		Response(c, http.StatusOK, gin.H{
			utils.ErrorCode: utils.ErrorCodeSuccess,
			utils.ErrorDesc: utils.ErrorDescSuccess,
			utils.Content:   content})
	}
}

func ResponseWithFailed(c *gin.Context, utilsErr *utils.Error) {
	Response(c, http.StatusOK, utilsErr)
}

func ResponseWithParameterError(c *gin.Context, errorDesc string) {
	Response(c, http.StatusBadRequest, utils.NewError(utils.ErrorCodeParameterError, errorDesc))
}

func ResponseWithMissingParameter(c *gin.Context, param string) {
	Response(c, http.StatusBadRequest, utils.NewErrorWithMissingParameter(param))
}
