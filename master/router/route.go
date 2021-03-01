package router

import (
	"easycron/master/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouters() (r *gin.Engine){

	r = gin.Default()

	// 全局中间件
	r.Use(middleware.MyFMT())

	// 路由模块化
	jobRouter(r)

	return
}
