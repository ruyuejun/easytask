package router

import (
	"ginserver/controller/demoCtr"
	"github.com/gin-gonic/gin"
)

func testRouter(router *gin.Engine) {

	router.GET("/demo", func(c *gin.Context){
		demoCtr.Demo(c)
	})

}
