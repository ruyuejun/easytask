package userModel

import (
	"ginserver/common/base"
)

// 账户接口规范
type accountI interface {
	Check() *base.Result    // 检测用户是否注册
	Register() *base.Result // 注册
	Login() *base.Result    // 登陆
	Logout() *base.Result   // 退出
}

// 用户接口规范
type userI interface {
	Auth() *base.Result // 鉴权
}
