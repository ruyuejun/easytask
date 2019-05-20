package userModel

import "ginserver/common"

type User struct {
	Uid int `json: "uid"`
	UserName string `json: "username"`
	PassWord string `json: "password"`
	Tel string `json:"tel"`
}

var code = &common.CODE

func (u *User)Find() *common.Code{
	return &common.Code{
	}
}
