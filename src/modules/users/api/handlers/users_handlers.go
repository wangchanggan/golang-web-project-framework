package handlers

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/src/modules/users/service"
	"golang-web-project-framework/src/modules/users/service/dto"
	"golang-web-project-framework/src/utils"
	utilsHttp "golang-web-project-framework/src/utils/http"
)

type UsersHandlers struct {
}

func (usersHandlers *UsersHandlers) AddUser(c *gin.Context) {
	var userAddReqDto dto.UserAddReqDto
	err := c.ShouldBindJSON(&userAddReqDto)
	if err != nil {
		utils.Warnf("c.ShouldBindJSON failed: error=%s", err.Error())
		utilsHttp.ResponseWithParameterError(c, err.Error())
		return
	}

	userAddRespDto, utilsErr := service.GetUserServiceInstance().AddUser(userAddReqDto)
	if utilsErr != nil {
		utils.Warnf("add User failed: error=%s", utilsErr.Error)
		utilsHttp.ResponseWithFailed(c, utilsErr)
		return
	}

	utilsHttp.ResponseWithOk(c, userAddRespDto)
}