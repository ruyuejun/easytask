package task

import (
	"fmt"
	"time"
)

// 任务调度器
type Scheduler struct {
	EventChan	chan *JobEvent           // 任务事件队列
	PlanTable	map[string]*SchedulePlan // 任务调度计划表
}

var GScheduler *Scheduler

func NewScheduler() {
	GScheduler = &Scheduler{
		EventChan: make(chan *JobEvent, 1000),
		PlanTable: make(map[string]*SchedulePlan),
	}
	go GScheduler.Loop()
}

// 调度协程
func (scheduler *Scheduler)Loop() {

	var scheduleAfter time.Duration
	var scheduleTimer *time.Timer
	var jobEvent *JobEvent

	// 初始化一次（1秒）
	scheduleAfter = scheduler.Try()
	// 调度的延时定时器
	scheduleTimer = time.NewTimer(scheduleAfter)

	// 定时任务
	for {
		select {
		case jobEvent = <-scheduler.EventChan:
			fmt.Println("取出的事件:", jobEvent.Job.Name)
			// 对内存中维护的任务列表做增删改查
			scheduler.Handle(jobEvent)
		case <-scheduleTimer.C:		// 等待最近的任务到期

		}
		// 立即调度一次任务，并重置时间间隔
		scheduleAfter = scheduler.Try()
		scheduleTimer.Reset(scheduleAfter)
	}
}

// 推送任务变化事件
func (scheduler *Scheduler)Push(event *JobEvent) {

	fmt.Println("函数内push的是：", event.Job)
	GScheduler.EventChan <- event

}

// 处理任务事件
func (scheduler *Scheduler)Handle(event *JobEvent) {

	var (
		schedulePlan *SchedulePlan
		jobExisted bool
		err error
	)

	switch event.EventType {

	case JOB_EVENT_SAVE:
		schedulePlan, err = NewSchedulePlan(event.Job)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		scheduler.PlanTable[event.Job.Name] = schedulePlan

	case JOB_EVENT_DELET:
		// 为了避免被反复删除的任务推入队列，先进行验证是否删除
		if schedulePlan, jobExisted =  scheduler.PlanTable[event.Job.Name]; jobExisted {
			delete(scheduler.PlanTable, event.Job.Name)
		}
	}
}

// 重新计算任务调度状态
func (scheduler *Scheduler)Try() (scheduleAfter time.Duration){

	var plan *SchedulePlan
	var nearTime *time.Time
	var nowTime time.Time

	// 如果任务表为空，睡眠任意时间
	if len(scheduler.PlanTable) == 0 {
		scheduleAfter = 1 * time.Second
		return
	}

	nowTime = time.Now()

	// 遍历所有任务
	for _, plan = range scheduler.PlanTable {

		if plan.NextTime.Before(nowTime) || plan.NextTime.Equal(nowTime) {
			// 尝试执行任务
			fmt.Println("执行任务：", plan.Job.Name)
			plan.NextTime = plan.Expr.Next(nowTime)		// 更新下次执行时间
		}

		// 统计最近一个要过期的任务时间（N秒后过期则为scheduleAfter）
		if nearTime == nil || plan.NextTime.Before(*nearTime){
			nearTime = &plan.NextTime
		}
	}

	// 下次调度间隔：最近要执行的任务调度时间-当前时间
	scheduleAfter = (*nearTime).Sub(nowTime)
	return
}
