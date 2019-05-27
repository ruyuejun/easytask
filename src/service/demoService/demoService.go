package demoService

import (
	"ginserver/common/code"
	"ginserver/model/demoModel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Demo(c *gin.Context) {

	demo := demoModel.Demo{
		"张三",
		"123456",
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.OK.Code,
		"msg": code.OK.Msg,
		"data": demo,
	})

}
