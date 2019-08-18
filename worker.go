package main

import (
	"dcs-gocron/config"
	"dcs-gocron/task"
	"flag"
	"fmt"
	"time"
)

func main() {

	// 解析命令行参数
	var confFile string
	flag.StringVar(&confFile, "config", "./config/config.json", "指定配置文件")
	flag.Parse()

	// 初始化配置文件
	err := config.NewConfig("./config/config.json")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// 启动执行器
	task.NewExecutor()

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
