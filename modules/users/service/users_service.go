package service

import (
	"golang-web-project-framework/modules/users/model"
	"sync"
	"golang-web-project-framework/modules/users/service/dto"
	utilsError "golang-web-project-framework/utils/error"
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

func (usersService *UsersService) AddUser(userAddReqDto dto.UserAddReqDto) (*dto.UserAddRespDto, *utilsError.Error) {
	return model.NewUser(userAddReqDto).AddUser()
}
