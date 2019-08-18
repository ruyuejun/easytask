package task

import (
	"github.com/gorhill/cronexpr"
	"time"
)

// 调度计划对象
type SchedulePlan struct {
	Job *Job                  // 要调度的任务
	Expr *cronexpr.Expression // 解析好的cronexpr表达式
	NextTime time.Time        // 任务下次执行时间
}

func NewSchedulePlan(job *Job) (schedulePlan *SchedulePlan, err error){

	// 表达式解析
	expr, err := cronexpr.Parse(job.CronExpr)
	if err != nil {
		return
	}

	// 生成任务调度计划对象
	schedulePlan = &SchedulePlan{
		Job:      job,
		Expr:     expr,
		NextTime: time.Now(),
	}
	return
}
