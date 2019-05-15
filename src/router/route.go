package router

import(
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	//路由模块化
	testRouter(router)				// 测试模块
	userRouter(router)				// 用户模块

	return router

}
