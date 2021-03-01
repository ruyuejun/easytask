package controller

import (
	"easycron/common"
	"easycron/master/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobController struct {}

func (ctr *JobController)SaveJob(c *gin.Context) {

	name := c.PostForm("name")
	command := c.PostForm("command")
	expr := c.PostForm("expr")

	var jobService *service.JobService
	resp := jobService.SaveJob(name, command, expr)

	common.BuildResponse(resp)
	c.JSON(http.StatusOK, resp)
}

func (ctr *JobController)DeleteJob(c *gin.Context) {

	name := c.PostForm("name")

	var jobService *service.JobService
	resp := jobService.DeleteJob(name)

	common.BuildResponse(resp)
	c.JSON(http.StatusOK, resp)
}

func (ctr *JobController)UpdateJob(c *gin.Context) {
	ctr.SaveJob(c)
}

func (ctr *JobController)ListJobs(c *gin.Context) {

	name := c.PostForm("name")

	var jobService *service.JobService
	resp := jobService.ListJobs(name)

	common.BuildResponse(resp)
	c.JSON(http.StatusOK, resp)
}


func (ctr *JobController)KillJob(c *gin.Context) {

	name := c.PostForm("name")

	var jobService *service.JobService
	resp := jobService.KillJob(name)

	common.BuildResponse(resp)
	c.JSON(http.StatusOK, resp)
}
