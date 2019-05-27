package redisUtil

import (
	"fmt"
	"github.com/go-redis/redis"
	"ginserver/config"
)

var conf =  &config.CONF

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {

	if RedisClient != nil {
		fmt.Println("RedisClient 不为空")
		return RedisClient
	}

	bf := config.InitBaseConfig()

	redisInfo := bf.REDIS_TCP + ":" + bf.REDIS_PORT

	client := redis.NewClient(&redis.Options{
		Addr:     redisInfo,
		Password: bf.REDIS_PASS,     // no password set
		DB:       bf.REDIS_DATABASE, // use default DB
	})

	//fmt.Println("type:", reflect.TypeOf(client))

	RedisClient = client

	return RedisClient
}
