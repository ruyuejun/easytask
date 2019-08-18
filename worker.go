package main

import (
	"dcs-gocron/config"
	"dcs-gocron/task"
	"fmt"
	"time"
)

func main() {

	// 初始化配置文件
	err := config.NewConfig("./config/config.json")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// 启动调度器
	task.NewScheduler()

	// 初始化任务管理器
	err = task.NewJobManager()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// 监听
	err = task.GJobManager.WatchJobs()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	for {
		time.Sleep(time.Second * 5)
	}
}
