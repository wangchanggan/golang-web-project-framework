package model

import (
	userDto "golang-web-project-framework/modules/users/service/dto"
	utilsError "golang-web-project-framework/utils/error"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	CreatedTime int64  `bson:"created_time"`
}

func NewUser(userAddReqDto userDto.UserAddReqDto) *User {
	return &User{
		Id:          bson.NewObjectId().Hex(),
		Name:        userAddReqDto.Name,
		CreatedTime: time.Now().Unix(),
	}
}

func (user *User) AddUser() (*userDto.UserAddRespDto, *utilsError.Error) {
	return nil, nil
}
