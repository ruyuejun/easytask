package main

import (
	"dcs-gocron/config"
	"dcs-gocron/task"
	"dcs-gocron/router"
	"fmt"
)

func main() {

	// 初始化配置文件
	err := config.NewConfig("./config/config.json")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// 初始化任务管理器
	err = task.NewJobManager()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//  初始化路由
	r := router.NewRouters()
	_ = r.Run(":" + config.GConfig.Port)

}
