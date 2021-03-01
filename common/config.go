package common

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiPort 			int 		`json:"apiPort"`
	ApiReadTimeout 		int 		`json:"apiReadTimeout"`
	ApiWriteTimeout 	int 		`json:"apiWriteTimeout"`
	WorkerEndPonts 		[]string 	`json:"workerEndPonts"`
	WorkerDialTimeout 	int 		`json:"workerDialTimeout"`
}

var GConfig *Config

func InitConfig(commonFilePath string){
	
	content, err := ioutil.ReadFile(commonFilePath)
	if err != nil {
		panic(err)
	}


	err = json.Unmarshal(content, &GConfig)
	if err != nil {
		panic(err)
	}
}
