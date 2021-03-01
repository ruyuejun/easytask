package main

import (
	"easycron/common"
	"easycron/master"
	"easycron/worker"
	"flag"
)

// choose server type, exp: go run main -s master\
func startServer(){
	var serverType string
	flag.StringVar(&serverType, "s", "master", "run master or worker")
	flag.Parse()
	switch serverType {
	case "master":
		master.Run()
	case "worker":
		worker.Run()
	}
}

func main()  {
	common.InitConfig("./config.json")
	startServer()
}
