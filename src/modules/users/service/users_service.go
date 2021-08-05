package service

import (
	"golang-web-project-framework/src/modules/users/model"
	"sync"
	"golang-web-project-framework/src/modules/users/service/dto"
	"golang-web-project-framework/src/utils"
)

var (
	instance *UsersService
	once sync.Once
)

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func GetUserServiceInstance() *UsersService {
	once.Do(func() {
		instance = NewUsersService()
	})
	return instance
}

func (usersService *UsersService) AddUser(userAddReqDto dto.UserAddReqDto) (*dto.UserAddRespDto, *utils.Error) {
	return model.NewUser(userAddReqDto).AddUser()
}
