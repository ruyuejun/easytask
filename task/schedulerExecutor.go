package task

import (
	"os/exec"
	"time"
)

// 任务执行器
type Executor struct {

}
var GExecutor *Executor

func NewExecutor() *Executor{
	return &Executor{}
}

// 任务执行结果
type ExecuteResult struct {
	ExecuteInfo *ScheduleExecuteInfo
	Output []byte
	Err error
	StartTime time.Time
	EndTime time.Time
}



// 执行任务
func (executor *Executor)ExecuteJob(info *ScheduleExecuteInfo) {
	go func() {

		startTime := time.Now()

		cmd := exec.Command("sh", "-c", info.Job.Command)
		output, err := cmd.Output()

		endTime := time.Now()

		// 任务执行完毕后，执行结果需要返回给scheduler以便从ExecutingTable中删除执行记录
		result := &ExecuteResult{
			ExecuteInfo: info,
			Output:      make([]byte, 0),
			Err:         err,
			StartTime:   startTime,
			EndTime:     endTime,
		}
		result.Output = output

		// 任务执行结果回传
		GScheduler.PushJobResult(result)

	}()
}
