package config

/*
	配置设置
 */

var ENV = 1 // 0 本地开发环境，1 线上测试环境，2 线上生产环境

// 服务器配置
type ServerConfig struct {
	ServerPort string
}

// MySql配置
type MysqlConfig struct {
	MysqlHost string
	MysqlPort string
	MysqlAuth string
	MysqlPass string
	MysqlDataBase string
}

// Redis配置
type RedisConfig struct {
	RedisHost string
	RedisPort string
	RedisAuth string
	RedisPass string
	RedisDataBase string
}

// 所有配置
type conf struct {
	ServerConfig
	MysqlConfig
	RedisConfig
}

// 导出一个配置实例
var CONF conf

func init() {
	switch ENV {
	case 0 :
		initDevConfig(&CONF)
	case 1 :
		initTestConfig(&CONF)
	case 2 :
		initProdConfig(&CONF)
	default:
		return
	}
}
