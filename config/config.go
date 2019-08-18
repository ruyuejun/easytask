package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	EtcdEndpoints []string `json:"etcd_endpoints"`
	Port string `json:"port"`
}

var GConfig *Config

func NewConfig(filename string) (err error){

	var conf Config

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// json反序列化
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return
	}
	GConfig = &conf
	return
}
