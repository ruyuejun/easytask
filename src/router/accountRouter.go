package router

import (
	"github.com/gin-gonic/gin"
)

func accountRouter(router *gin.Engine) {

	// 账户组
	account := router.Group("/account")
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

}
