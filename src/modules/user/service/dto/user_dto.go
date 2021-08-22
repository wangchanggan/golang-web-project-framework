package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-web-project-framework/src/modules/user/model"
	utilsHttp "golang-web-project-framework/src/utils/http"
	"time"
)

type UserAddReqDto struct {
	Name string `json:"name" binding:"required"`
}

type UserUpdateReqDto struct {
	Name string `json:"name" binding:"required"`
}

type UserAddRespDto struct {
	Id string `json:"id"`
}

type UsersRespDto struct {
	PageInfo utilsHttp.PageInfo `json:"pageInfo"`
	UserList []UserRespDto      `json:"userList"`
}

type UserRespDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CreatedTime int64  `json:"createdTime"`
}

type UserDetailRespDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CreatedTime int64  `json:"createdTime"`
}

func NewAddUser(userAddReqDto UserAddReqDto) model.UserDetailInfo {
	return model.UserDetailInfo{
		Id:          primitive.NewObjectID().Hex(),
		Name:        userAddReqDto.Name,
		CreatedTime: time.Now().Unix(),
	}
}

func NewUserUpdate(userUpdateReqDto UserUpdateReqDto) model.UserUpdate {
	return model.UserUpdate{
		Name: userUpdateReqDto.Name,
	}
}

func GenerateUsersRespDto(users []model.User, pageSize, totalSize int64) (usersRespDto *UsersRespDto) {
	userList := make([]UserRespDto, len(users))
	for k, v := range users {
		var tmp UserRespDto
		tmp.Id = v.Id
		tmp.Name = v.Name
		tmp.CreatedTime = v.CreatedTime
		userList[k] = tmp
	}

	usersRespDto = new(UsersRespDto)
	usersRespDto.UserList = userList
	usersRespDto.PageInfo = utilsHttp.GeneratePageInfo(pageSize, totalSize)
	return
}

func GenerateUserDetailRespDto(userDetailInfo *model.UserDetailInfo) *UserDetailRespDto {
	return &UserDetailRespDto{
		Id:          userDetailInfo.Id,
		Name:        userDetailInfo.Name,
		CreatedTime: userDetailInfo.CreatedTime,
	}
}
