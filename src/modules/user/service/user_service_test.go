package service

import (
	"fmt"
	"golang-web-project-framework/src/modules/user/service/dto"
	"testing"
)

func TestAddUser(t *testing.T) {
	var userAddReqDto dto.UserAddReqDto
	userAddReqDto.Name = "zhangsan"
	fmt.Println(GetUserServiceInstance().AddUser(userAddReqDto))
}

func TestUpdateUserById(t *testing.T) {
	var userUpdateReqDto dto.UserUpdateReqDto
	userUpdateReqDto.Name = "lisi"
	fmt.Println(GetUserServiceInstance().UpdateUserById("611a803f50059bd10934ecf1", userUpdateReqDto))
}

func TestGetUsers(t *testing.T) {
	usersRespDto, utilsErr := GetUserServiceInstance().GetUsers(1, 10, "created_time", "desc", "", "")
	if utilsErr != nil {
		fmt.Println(utilsErr)
	}else {
		fmt.Println(fmt.Sprintf("%+v", usersRespDto))
	}
}

func TestGetUser(t *testing.T) {
	userDetailRespDto, utilsErr := GetUserServiceInstance().GetUserById("611a803f50059bd10934ecf1")
	if utilsErr != nil {
		fmt.Println(utilsErr)
	}else {
		fmt.Println(fmt.Sprintf("%+v", userDetailRespDto))
	}
}

func TestDeleteUserById(t *testing.T) {
	fmt.Println(GetUserServiceInstance().DeleteUserById("611a803f50059bd10934ecf1"))
}
