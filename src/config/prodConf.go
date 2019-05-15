package config

/*
	线上生产环境
 */

func initProdConfig(c *conf) {

	c.ServerConfig = ServerConfig{
		"3001",
	}

	c.MysqlConfig = MysqlConfig{
		"",
		"",
		"",
		"",
		"",
	}

	c.RedisConfig = RedisConfig{
		"",
		"",
		"",
		"",
		"",
	}

}
