package task

type Job struct {
	Name string `json:"name"`				// 任务名
	Command string `json:"command"`			// 任务命令
	Expr string `json:"expr"`				// 任务cron
}