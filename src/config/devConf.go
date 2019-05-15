package config

/*
	本地开发环境
*/

func initDevConfig(c *conf) {

	c.ServerConfig = ServerConfig{
		"",
	}

	c.MysqlConfig = MysqlConfig{
		"localhost",
		"3306",
		"root",
		"",
		"",
	}

	c.RedisConfig = RedisConfig{
		"localhost",
		"6379",
		"",
		"",
		"9",
	}

}
