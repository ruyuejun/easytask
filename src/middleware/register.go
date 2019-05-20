package middleware

import (
	"ginserver/common"
	"ginserver/model/userModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

var code = &common.CODE

func isRegister() gin.HandlerFunc{

	return func(c *gin.Context) {

		tel := c.PostForm("tel")

		// 查询用户是否注册
		var u userModel.User
		u.Tel = tel
		r := u.Find()

		// 服务器错误
		if r.Code > 3000 {
			c.JSON()
		}

		// 用户未注册
		if r.Code != 2001 {
			c.JSON(http.StatusOK, gin.H{
				"code": code.OK.Code,
				"msg": code.OK.Msg,
				"data": nil,
			})
		}

		// 用户已注册
		c.Next()
	}

}
