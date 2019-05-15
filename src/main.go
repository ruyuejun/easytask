package main

import (
	"ginserver/config"
	"ginserver/router"
)

func init() {

	conf := config.CONF

	//初始化路由
	router.InitRouter().Run(":" + conf.ServerPort)

}

func main() {


}
