package demoService

import (
	"ginserver/config"
	"ginserver/model/demoModel"
	"net/http"

	"github.com/gin-gonic/gin"
)

var code = &config.CODE

func Demo(c *gin.Context) {

	demo := demoModel.Demo{
		"张三",
		"123456",
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.OK.Code,
		"msg": code.OK.Msg,
		"data": demo,
	})

	//
	//db := mysqlUtil.NewMySqlClient()
	//
	//sql := "SELECT * FROM user"
	//
	//results, err := db.QueryString(sql)
	//
	//logUtil.Log.Info("HelloHAHAH0000")
	//
	//if err != nil {
	//	logUtil.Log.Info("HelloHAHAH")
	//	fmt.Println("查询数据库出错：", err)
	//	panic(err)
	//}
	//
	//RedisClient := redisUtil.NewRedisClient()
	//
	//v, e := RedisClient.Get("ping").Result()
	//if e != nil {
	//	panic(e)
	//}
	//
	//fmt.Println("redis v ==", v)
	//
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code": config.InitCodeConfig()["OK"].Code,
	//	"msg": config.InitCodeConfig()["OK"].Msg,
	//	"data": results,
	//})

}
