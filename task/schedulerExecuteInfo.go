package task

import "time"

// 任务执行状态
type ScheduleExecuteInfo struct {
	Job *Job
	PlanTime time.Time 		// 理论调度时间
	RealTime time.Time		// 实际调度时间
}

func NewScheduleExecuteInfo(plan *SchedulePlan) *ScheduleExecuteInfo{
	return &ScheduleExecuteInfo{
		Job:      plan.Job,
		PlanTime: plan.NextTime,
		RealTime: time.Now(),
	}
}
