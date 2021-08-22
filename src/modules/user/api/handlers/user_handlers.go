package handlers

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/src/modules/user/service"
	"golang-web-project-framework/src/modules/user/service/dto"
	"golang-web-project-framework/src/utils"
	utilsHttp "golang-web-project-framework/src/utils/http"
)

const UserId = "userId"

type UserHandlers struct {
}

func (userHandlers *UserHandlers) AddUser(c *gin.Context) {
	var userAddReqDto dto.UserAddReqDto
	err := utilsHttp.GetParamsFromBodyByJson(c, &userAddReqDto)
	if err != nil {
		utils.Warnf("c.ShouldBindJSON failed: error=%s", err.Error())
		utilsHttp.ResponseWithParameterError(c, err.Error())
		return
	}

	userAddRespDto, utilsErr := service.GetUserServiceInstance().AddUser(userAddReqDto)
	if utilsErr != nil {
		utils.Warnf("add User failed: error=%s", utilsErr.ErrorDesc)
		utilsHttp.ResponseWithFailed(c, utilsErr)
		return
	}

	utilsHttp.ResponseWithOk(c, userAddRespDto)
}

func (userHandlers *UserHandlers) DeleteUserById(c *gin.Context) {
	userId := utilsHttp.GetParamFromUrl(c, UserId)
	if userId == "" {
		utilsHttp.ResponseWithMissingParameter(c, UserId)
		return
	}

	utilsErr := service.GetUserServiceInstance().DeleteUserById(userId)
	if utilsErr != nil {
		utils.Warnf("delete User failed: error=%s", utilsErr.ErrorDesc)
		utilsHttp.ResponseWithFailed(c, utilsErr)
		return
	}

	utilsHttp.ResponseWithOk(c, nil)
}

func (userHandlers *UserHandlers) UpdateUserById(c *gin.Context) {
	userId := utilsHttp.GetParamFromUrl(c, UserId)
	if userId == "" {
		utilsHttp.ResponseWithMissingParameter(c, UserId)
		return
	}

	var userUpdateReqDto dto.UserUpdateReqDto
	err := utilsHttp.GetParamsFromBodyByJson(c, &userUpdateReqDto)
	if err != nil {
		utils.Warnf("c.ShouldBindJSON failed: error=%s", err.Error())
		utilsHttp.ResponseWithParameterError(c, err.Error())
		return
	}

	utilsErr := service.GetUserServiceInstance().UpdateUserById(userId, userUpdateReqDto)
	if utilsErr != nil {
		utils.Warnf("update User failed: error=%s", utilsErr.ErrorDesc)
		utilsHttp.ResponseWithFailed(c, utilsErr)
		return
	}

	utilsHttp.ResponseWithOk(c, nil)
}

func (userHandlers *UserHandlers) GetUsers(c *gin.Context) {
	pageNum, pageSize, order, orderBy, searchValue, searchBy, err := utilsHttp.GenerateListInfo(c)
	if err != nil {
		utils.Warnf("generate List Info failed: error=%s", err.Error())
		utilsHttp.ResponseWithParameterError(c, err.Error())
		return
	}

	result, utilsErr := service.GetUserServiceInstance().GetUsers(pageNum, pageSize, order, orderBy, searchValue, searchBy)
	if utilsErr != nil {
		utils.Warnf("get users failed: error=%s", utilsErr.ErrorDesc)
		utilsHttp.ResponseWithFailed(c, utilsErr)
		return
	}

	utilsHttp.ResponseWithOk(c, result)
}

func (userHandlers *UserHandlers) GetUserById(c *gin.Context) {
	userId := utilsHttp.GetParamFromUrl(c, UserId)
	if userId == "" {
		utilsHttp.ResponseWithMissingParameter(c, UserId)
		return
	}

	result, utilsErr := service.GetUserServiceInstance().GetUserById(userId)
	if utilsErr != nil {
		utils.Warnf("get User failed: error=%s", utilsErr.ErrorDesc)
		utilsHttp.ResponseWithFailed(c, utilsErr)
		return
	}

	utilsHttp.ResponseWithOk(c, result)
}
