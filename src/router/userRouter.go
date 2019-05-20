package router

import (
	"ginserver/controller/userCtr"
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.Engine) {

	// 账户组
	account := router.Group("/user")
	{
		// 登录：以手机为核心
		account.POST("/signin", func(context *gin.Context) {
			userCtr.SignIn(context)
		})

		// 退出
		account.POST("/signout", func(context *gin.Context) {
			userCtr.SignOut(context)
		})

	}

	// 个人信息组
	profile := router.Group("/user")
	{
		// 修改个人信息
		profile.PUT("/profile", func(context *gin.Context) {

		})

		// 获取个人信息
		profile.GET("/profile", func(context *gin.Context) {

		})

	}

}
