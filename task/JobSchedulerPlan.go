package task

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

// 某个任务的调度计划
type JobSchedulePlan struct {
	CurrentJob 		*Job
	CurrentExpr 	*cronexpr.Expression
	NextTime 		time.Time	// 下次调度时间
}

func NewJobSchedulerPlan(job *Job)(jobSchedulePlan *JobSchedulePlan, err error){
	fmt.Println("new job plan ：", job.Expr)
	expr, err := cronexpr.Parse(job.Expr)
	if err != nil {
		fmt.Println("new JobSchedulerPlan err:", err)
		return
	}

	jobSchedulePlan = &JobSchedulePlan{
		CurrentJob: 	job,
		CurrentExpr: 	expr,
		NextTime: 		expr.Next(time.Now()),	
	}

	return
}