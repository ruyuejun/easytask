package masterCtr

import (
	"dcs-gocron/common"
	"dcs-gocron/task"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveJob(c *gin.Context) {

	name := c.PostForm("name")
	command := c.PostForm("command")
	cronExpr := c.PostForm("cronExpr")

	job := &task.Job{
		Name:    name,
		Command: command,
		CronExpr: cronExpr,
	}

	oldJob, err := task.GJobManager.SaveJob(job)
	if err != nil {
		c.JSON(http.StatusOK, common.NewCode(common.ServerErr, err))
	}
	c.JSON(http.StatusOK, common.NewCode(common.OK, oldJob))
}

func DeleteJob(c *gin.Context) {

	name := c.PostForm("name")
	command := c.PostForm("command")
	cronExpr := c.PostForm("cronExpr")

	job := &task.Job{
		Name:    name,
		Command: command,
		CronExpr: cronExpr,
	}

	oldJob, err := task.GJobManager.DeleteJob(job)
	if err != nil {
		c.JSON(http.StatusOK, common.NewCode(common.ServerErr, err))
	}
	c.JSON(http.StatusOK, common.NewCode(common.OK, oldJob))
}

func KillJob(c *gin.Context) {

	name := c.PostForm("name")
	command := c.PostForm("command")
	cronExpr := c.PostForm("cronExpr")

	job := &task.Job{
		Name:    name,
		Command: command,
		CronExpr: cronExpr,
	}

	oldJob, err := task.GJobManager.KillJob(job)
	if err != nil {
		c.JSON(http.StatusOK, common.NewCode(common.ServerErr, err))
	}
	c.JSON(http.StatusOK, common.NewCode(common.OK, oldJob))
}

func ListJobs(c *gin.Context) {

	jobList, err := task.GJobManager.ListJobs()
	if err != nil {
		c.JSON(http.StatusOK, common.NewCode(common.ServerErr, err))
	}
	c.JSON(http.StatusOK, common.NewCode(common.OK, jobList))
}
