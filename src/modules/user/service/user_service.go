package service

import (
	"golang-web-project-framework/src/modules/user/model"
	"sync"
	"golang-web-project-framework/src/modules/user/service/dto"
	"golang-web-project-framework/src/utils"
	utilsHttp "golang-web-project-framework/src/utils/http"
)

var (
	userService *UserService
	once sync.Once
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func GetUserServiceInstance() *UserService {
	once.Do(func() {
		userService = NewUserService()
	})
	return userService
}

func (userService *UserService) AddUser(userAddReqDto dto.UserAddReqDto) (*dto.UserAddRespDto, *utils.Error) {
	id, utilsErr :=  model.AddUser(dto.NewAddUser(userAddReqDto))
	if utilsErr != nil {
		utils.Warnf(" add user failed: error=%s", utilsErr.ErrorDesc)
		return nil, utilsErr
	}

	return &dto.UserAddRespDto{
		Id: id,
	}, nil
}

func (userService *UserService) DeleteUserById(userId string) *utils.Error {
	return model.DeleteUserById(userId)
}

func (userService *UserService) UpdateUserById(userId string, userUpdateReqDto dto.UserUpdateReqDto) *utils.Error {
	return model.UpdateUserById(userId, dto.NewUserUpdate(userUpdateReqDto))
}

func (userService *UserService) GetUsers(pageNum, pageSize int64, order, orderBy, searchValue, searchBy string) (usersRespDto *dto.UsersRespDto, utilsErr *utils.Error) {
	pageNum, pageSize, order, orderBy = utilsHttp.CompleteListInfo(pageNum, pageSize, order, orderBy)
	invalidParam := utilsHttp.ValidateListInfo(searchValue, searchBy)
	if invalidParam != "" {
		return nil, utils.NewErrorWithInvalidParameter(invalidParam)
	}

	users, totalSize, utilsErr := model.GetUsers(pageNum, pageSize, order, orderBy, searchValue, searchBy)
	if utilsErr != nil {
		utils.Warnf(" get users failed: error=%s", utilsErr.ErrorDesc)
		return
	}

	return dto.GenerateUsersRespDto(users, pageSize, totalSize), nil
}

func (userService *UserService) GetUserById(userId string) (userDetailRespDto *dto.UserDetailRespDto, utilsErr *utils.Error) {
	user, utilsErr := model.GetUserById(userId)
	if utilsErr != nil {
		utils.Warnf(" get user by id failed: userId=%s, error=%s", userId, utilsErr.ErrorDesc)
		return
	}

	return dto.GenerateUserDetailRespDto(user), nil
}

