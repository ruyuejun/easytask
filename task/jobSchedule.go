package task

import (
	"fmt"
	"time"
)

type JobScheduler struct {
	jobEventChan chan *JobEvent
	jobPlanTable map[string]*JobSchedulePlan	// 任务调度计划表
}

var GScheduler *JobScheduler

// 初始化调度协程
func InitScheduler(){
	GScheduler = &JobScheduler{
		jobEventChan: make(chan *JobEvent, 1000),
		jobPlanTable: make(map[string] *JobSchedulePlan),
	}
	go GScheduler.scheduleLoop()
}

// 重复循环监听任务变化事件，一旦有任务，则对内存中维护的任务进行同步
func (scheduler *JobScheduler) scheduleLoop(){

	// 启动前先计算一次
	schedulerAfter := scheduler.computedSchedule()
	schedulerTimer := time.NewTimer(schedulerAfter)

	var jobEvent *JobEvent
	for {	
		select {
		case jobEvent = <-scheduler.jobEventChan:// 监听任务事件变化
			scheduler.handleJobEvent(jobEvent)
		case <- schedulerTimer.C:// 最近任务到期
		}
		scheduleAfter := scheduler.computedSchedule()
		schedulerTimer.Reset(scheduleAfter)
	}
}

// 推送任务变化事件
func (scheduler *JobScheduler) PushJobEvent(jobEvent *JobEvent){
	scheduler.jobEventChan <- jobEvent
}

// 处理任务事件
func (scheduler *JobScheduler) handleJobEvent(jobEvent *JobEvent){
	switch jobEvent.EventType {
	case GC.JOB_EVENT_SAVE:
		jobSchedulePlan, err := NewJobSchedulerPlan(jobEvent.EventJob)
		if err != nil {
			return
		}
		scheduler.jobPlanTable[jobEvent.EventJob.Name] = jobSchedulePlan

	case GC.JOB_EVENT_DELETE:
		// 先判断任务是否仍然存在
		_, isExisted := scheduler.jobPlanTable[jobEvent.EventJob.Name]
		if isExisted == true {
			delete(scheduler.jobPlanTable, jobEvent.EventJob.Name)
		}
	}
}

//  优化调度任务状态
func (scheduler *JobScheduler) computedSchedule() (scheduleAfter time.Duration){

	var nearTime *time.Time
	now := time.Now()

	// 如果任务表为空
	if len(scheduler.jobPlanTable) == 0 {
		scheduleAfter = 1 * time.Second
		return
	}

	// 遍历所有任务
	for _,jobPlan := range scheduler.jobPlanTable {
		if jobPlan.NextTime.Before(now) || jobPlan.NextTime.Equal(now) {
			// 尝试执行任务
			fmt.Println("执行任务：", jobPlan.CurrentJob.Name)
			jobPlan.NextTime = jobPlan.CurrentExpr.Next(now)
		}

		// 计算最近要过期的任务事件
		if nearTime == nil || jobPlan.NextTime.Before(*nearTime) {
			nearTime = &jobPlan.NextTime
		}
	}

	// 下次调度间隔（最近要执行的任务调度时间 - 当前时间）
	scheduleAfter = (*nearTime).Sub(now)
	return
}
