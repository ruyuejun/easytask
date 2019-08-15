package masterCtr

import (
	"Demo1/common"
	"Demo1/jobmanager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveJob(c *gin.Context) {

	name := c.PostForm("name")
	command := c.PostForm("command")
	cronExpr := c.PostForm("cronExpr")

	job := &jobmanager.Job{
		Name:    name,
		Command: command,
		CronExpr: cronExpr,
	}

	oldJob, err := jobmanager.GJobManager.SaveJob(job)
	if err != nil {
		c.JSON(http.StatusOK, common.CreateCode(common.ServerErr, err))
	}
	c.JSON(http.StatusOK, common.CreateCode(common.OK, oldJob))
}

func DeleteJob(c *gin.Context) {

	name := c.PostForm("name")
	command := c.PostForm("command")
	cronExpr := c.PostForm("cronExpr")

	job := &jobmanager.Job{
		Name:    name,
		Command: command,
		CronExpr: cronExpr,
	}

	oldJob, err := jobmanager.GJobManager.DeleteJob(job)
	if err != nil {
		c.JSON(http.StatusOK, common.CreateCode(common.ServerErr, err))
	}
	c.JSON(http.StatusOK, common.CreateCode(common.OK, oldJob))
}

func KillJob(c *gin.Context) {

	name := c.PostForm("name")
	command := c.PostForm("command")
	cronExpr := c.PostForm("cronExpr")

	job := &jobmanager.Job{
		Name:    name,
		Command: command,
		CronExpr: cronExpr,
	}

	oldJob, err := jobmanager.GJobManager.KillJob(job)
	if err != nil {
		c.JSON(http.StatusOK, common.CreateCode(common.ServerErr, err))
	}
	c.JSON(http.StatusOK, common.CreateCode(common.OK, oldJob))
}

func ListJobs(c *gin.Context) {

	jobList, err := jobmanager.GJobManager.ListJobs()
	if err != nil {
		c.JSON(http.StatusOK, common.CreateCode(common.ServerErr, err))
	}
	c.JSON(http.StatusOK, common.CreateCode(common.OK, jobList))
}
