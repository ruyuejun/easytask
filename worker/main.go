package worker

import "easycron/task"

func Run(){

	// 初始化执行器
	task.InitExecutor()

	// 初始化调度器
	task.InitScheduler()

	// 初始化管理者
	task.InitManager()

	for{
		
	}
}
