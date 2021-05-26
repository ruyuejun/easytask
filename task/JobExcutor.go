package task

import (
	"context"
	"os/exec"
	"time"
)

// 任务执行器
type JobExecutor struct{

}

// 任务执行结果
type JobExecutorRes struct{
	Executer	*JobSchedulerExcuter
	Output		[]byte			// 脚本输出
	Error 		error			// 脚本错误
	StartTime	time.Time		// 启动时间
	EndTime		time.Time		// 结束时间
}

var GE *JobExecutor

// 初始化执行器
func InitExecutor(){
	GE = &JobExecutor{}
}

func (executor *JobExecutor)execute(jobExcuter *JobSchedulerExcuter) (res *JobExecutorRes){
	// 执行 shell 命令
	starTime := time.Now()
	go func(){
		cmd := exec.CommandContext(context.TODO(), "/bin/bash", "-c", jobExcuter.excuteJob.Command)
		output, err := cmd.CombinedOutput()
		// 返回执行结果，让Scheduler 从 执行表 中删除执行记录
		res.Executer = jobExcuter
		res.Output = output
		res.Error = err
		res.StartTime = starTime
		res.EndTime = time.Now()
	}()
	return
}