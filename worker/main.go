package worker

import "easycron/task"

func Run(){
	// 初始化管理者
	task.InitManager()

	// 初始化调度器
	task.InitScheduler()

	// 启动监视
	task.GM.Watch()

	for{
		
	}
}
