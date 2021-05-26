package task

import "time"

type JobSchedulerExcuter struct{
	excuteJob 	*Job		
	PlanTime 	time.Time	// 理论调度时间
	RealTime	time.Time	// 实际调度时间
}

func NewJobSchedulerExcuter(plan *JobSchedulePlan) *JobSchedulerExcuter{
	return &JobSchedulerExcuter{
		excuteJob: plan.CurrentJob,
		PlanTime: plan.NextTime,
		RealTime: time.Now(),
	}
}