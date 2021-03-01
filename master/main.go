package master

import (
	"easycron/common"
	"easycron/master/router"
	"easycron/task"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Run(){
	task.InitManager()

	gin.SetMode(gin.DebugMode)
	gin.ForceConsoleColor()
	
	router.NewRouters().Run(":" + strconv.Itoa(common.GConfig.ApiPort))
}
