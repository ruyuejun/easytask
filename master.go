package main

import (
	"Demo1/config"
	"Demo1/jobmanager"
	"Demo1/router"
	"fmt"
)

func main() {

	// 初始化配置文件
	err := config.Init("./config/config.json")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// 初始化etcd连接
	err = jobmanager.Init()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//  初始化路由
	r := router.InitRouter()
	_ = r.Run(":" + config.GConfig.Port)

}
