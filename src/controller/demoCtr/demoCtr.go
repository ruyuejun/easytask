package demoCtr

import (
	"ginserver/service/demoService"
	"github.com/gin-gonic/gin"
)

func Demo(c *gin.Context) {

	//参数验证

	//具体的service
	demoService.Demo(c)

}
