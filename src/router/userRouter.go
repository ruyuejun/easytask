package router

import (
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.Engine) {

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
