package router

import (
	"dcs-gocron/web/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	// 全局中间件
	r.Use(middleware.MyFMT())

	// 路由模块化
	masterRouter(r)

	return r
}
