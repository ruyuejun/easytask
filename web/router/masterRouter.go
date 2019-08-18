package router

import (
	"dcs-gocron/web/controllers/masterCtr"
	"github.com/gin-gonic/gin"
)

func masterRouter(r *gin.Engine) {

	// 保存任务
	r.POST("/job/save", masterCtr.SaveJob)

	// 删除任务
	r.POST("job/delete", masterCtr.DeleteJob)

	// 杀死任务
	r.POST("job/kill", masterCtr.KillJob)

	// 查询任务
	r.POST("job/list", masterCtr.ListJobs)

}
