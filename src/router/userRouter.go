package router

import (
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.Engine) {

	// 账户组
	account := router.Group("/user/account")
	{

		// 用户注册
		account.POST("/register", func(context *gin.Context) {

		})

		// 用户登录
		account.POST("/login", func(context *gin.Context) {

		})

		// 用户退出
		account.POST("/logout", func(context *gin.Context) {

		})

	}

	// 个人信息组
	profile := router.Group("/user/profile")
	{
		// 修改个人信息
		profile.PUT("", func(context *gin.Context) {

		})

		// 获取个人信息
		profile.GET("", func(context *gin.Context) {
			
		})

	}

}
