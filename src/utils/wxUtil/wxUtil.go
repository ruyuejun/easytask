package wxUtil

import (
	"fmt"
	"ginserver/config"
)

var conf = &config.CONF

func GetCode() {
	fmt.Println(conf)
}
