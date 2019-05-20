package userModel

import (
	"ginserver/common"
)

// 账户接口规范
type accountI interface {
	Check() *common.Result					// 检测用户是否注册
	Register() *common.Result				// 注册
	Login() *common.Result					// 登陆
	Logout() *common.Result					// 退出
}

// 用户接口规范
type userI interface {
	Auth() *common.Result					// 鉴权
}
