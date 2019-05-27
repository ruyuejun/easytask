package mysqlUtil

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"petapi/config"
)

var SQLClient *xorm.Engine

func NewMySqlClient() *xorm.Engine{

	if SQLClient != nil {
		fmt.Println("SQLClient 不为空")
		return SQLClient
	}

	fmt.Println("SQLClient 为空")

	var bf = config.InitBaseConfig()

	MySqlInfo := bf.MYSQL_USER + ":" + bf.MYSQL_PASS + "@tcp(" + bf.MYSQL_TCP + ":" + bf.MYSQL_PORT + ")/" + bf.MYSQL_DATABASE + "?charset=utf8"

	SQLClient, err := xorm.NewEngine("mysql", MySqlInfo)

	if err != nil {
		fmt.Println("xorm 启动错误：", err)
		panic(err)
	}

	if err := SQLClient.Ping(); err != nil {
		fmt.Println("xorm 连接错误：", err)
		panic(err)
	}

	//日志打印
	SQLClient.ShowSQL(true)

	//设置连接池的空闲数大小
	SQLClient.SetMaxIdleConns(5)
	//设置最大打开连接数
	SQLClient.SetMaxOpenConns(5)

	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	SQLClient.SetTableMapper(core.SnakeMapper{})

	//fmt.Println("type:", reflect.TypeOf(engine))

	return SQLClient

}