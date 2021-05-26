package task

import (
	"fmt"
	"time"
)

// 调度协程
type JobScheduler struct {
	jobEventChan chan *JobEvent
	jobPlanTable map[string]*JobSchedulePlan			// 任务调度表：用来发现哪个任务过期
	jobExcutingTable map[string]*JobSchedulerExcuter	// 任务执行表：所有执行中的任务，用于优化一些执行时间较长的任务
	jobResultChan chan *JobExecutorRes	// 任务结果队列
}

var GScheduler *JobScheduler

// 初始化调度协程
func InitScheduler(){
	GScheduler = &JobScheduler{
		jobEventChan: make(chan *JobEvent, 1000),
		jobPlanTable: make(map[string] *JobSchedulePlan),
		jobExcutingTable: make(map[string]*JobSchedulerExcuter),
	}

	// 启动调度协程：发现哪个任务过期
	go func ()  {
		// 启动前先计算一次
		schedulerAfter := GScheduler.computedSchedule()
		schedulerTimer := time.NewTimer(schedulerAfter)

		// 循环监听任务变化事件，一旦有任务，则对内存中维护的任务进行同步
		var jobEvent *JobEvent
		for {	
			select {
			case jobEvent = <-GScheduler.jobEventChan:// 监听任务事件变化
				GScheduler.handleJobEvent(jobEvent)
			
			case <- schedulerTimer.C:// 最近任务到期
			
			case jobExecRes := <-GScheduler.jobResultChan:
				GScheduler.handleJobResult(jobExecRes)
			}

			// 调度一次任务
			scheduleAfter := GScheduler.computedSchedule()
			// 重置调度间隔
			schedulerTimer.Reset(scheduleAfter)
		}
	}()
}

// 推送任务变化事件
func (scheduler *JobScheduler) PushJobEvent(jobEvent *JobEvent){
	scheduler.jobEventChan <- jobEvent
}

// 在内存中处理任务事件，以保持内存与etcd一致
func (scheduler *JobScheduler) handleJobEvent(jobEvent *JobEvent){
	switch jobEvent.EventType {
	case GC.JOB_EVENT_SAVE:	// 保存任务事件
		jobSchedulePlan, err := NewJobSchedulerPlan(jobEvent.EventJob)
		if err != nil {
			return
		}
		scheduler.jobPlanTable[jobEvent.EventJob.Name] = jobSchedulePlan

	case GC.JOB_EVENT_DELETE:	// 删除任务事件
		// 先判断任务是否仍然存在
		_, isExisted := scheduler.jobPlanTable[jobEvent.EventJob.Name]
		if isExisted == true {
			delete(scheduler.jobPlanTable, jobEvent.EventJob.Name)
		}
	}
}

func (scheduler *JobScheduler) handleJobResult(execRes *JobExecutorRes){

		// 删除执行状态
		delete(scheduler.jobExcutingTable, execRes.Executer.excuteJob.Name)
	
		// 生成执行日志
		// var (
		// 	jobLog *common.JobLog
		// )
		// if result.Err != common.ERR_LOCK_ALREADY_REQUIRED {
		// 	jobLog = &common.JobLog{
		// 		JobName: result.ExecuteInfo.Job.Name,
		// 		Command: result.ExecuteInfo.Job.Command,
		// 		Output: string(result.Output),
		// 		PlanTime: result.ExecuteInfo.PlanTime.UnixNano() / 1000 / 1000,
		// 		ScheduleTime: result.ExecuteInfo.RealTime.UnixNano() / 1000 / 1000,
		// 		StartTime: result.StartTime.UnixNano() / 1000 / 1000,
		// 		EndTime: result.EndTime.UnixNano() / 1000 / 1000,
		// 	}
		// 	if result.Err != nil {
		// 		jobLog.Err = result.Err.Error()
		// 	} else {
		// 		jobLog.Err = ""
		// 	}
		// 	G_logSink.Append(jobLog)
		// }
	
		// fmt.Println("任务执行完成:", result.ExecuteInfo.Job.Name, string(result.Output), result.Err)
}

//  优化调度任务状态：遍历所有任务，过期任务立即执行，统计最近要过期任务的时间
func (scheduler *JobScheduler) computedSchedule() (scheduleAfter time.Duration){

	var nearTime *time.Time
	now := time.Now()

	// 如果任务表为空：睡眠1秒
	if len(scheduler.jobPlanTable) == 0 {
		scheduleAfter = 1 * time.Second
		return
	}

	// 遍历所有任务
	for _,jobPlan := range scheduler.jobPlanTable {
		if jobPlan.NextTime.Before(now) || jobPlan.NextTime.Equal(now) {
			// 尝试执行任务
			scheduler.startJob(jobPlan)
			fmt.Println("执行任务：", jobPlan.CurrentJob.Name)
			jobPlan.NextTime = jobPlan.CurrentExpr.Next(now)
		}

		// 计算最近要过期的任务时间
		if nearTime == nil || jobPlan.NextTime.Before(*nearTime) {
			nearTime = &jobPlan.NextTime
		}
	}

	// 下次调度间隔（最近要执行的任务调度时间 - 当前时间）
	scheduleAfter = (*nearTime).Sub(now)
	return
}

// 过期任务执行
func (scheduler *JobScheduler) startJob(plan *JobSchedulePlan){
	// 优化问题：一些任务1分钟调度多次，但是只会执行1次
	flag := scheduler.jobExcutingTable[plan.CurrentJob.Name]
	// 如果正在执行，则跳过本次调度
	if flag != nil {
		fmt.Println("跳过任务：", plan.CurrentJob.Name)
		return
	}

	// 不存在：则放入执行表
	excuteJob := NewJobSchedulerExcuter(plan)
	scheduler.jobExcutingTable[plan.CurrentJob.Name] = excuteJob
	
	// 执行任务：
	fmt.Println("执行任务", plan.CurrentJob.Name)
	GE.execute(excuteJob)
}