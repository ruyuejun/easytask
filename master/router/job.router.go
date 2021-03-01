package router

import (
	"easycron/master/controller"

	"github.com/gin-gonic/gin"
)

var jobCtr *controller.JobController

func jobRouter(r *gin.Engine) {

	// 新增任务
	r.POST("/job", jobCtr.SaveJob)

	// 删除任务
	r.DELETE("/job", jobCtr.DeleteJob)

	// 修改任务
	r.PUT("/job", jobCtr.UpdateJob)

	// 查询任务列表
	r.GET("/job", jobCtr.ListJobs)

	// 杀死任务
	r.POST("/job/kill", jobCtr.KillJob)

}
