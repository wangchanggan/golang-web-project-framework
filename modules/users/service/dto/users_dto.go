package dto

type UserAddReqDto struct {
	Name       string `json:"name" binding:"required"`
}

type UserAddRespDto struct {
	Id string `json:"id"`
}

type UserDetailRespDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday,omitempty"`
	Age      string `json:"age,omitempty"`
}
