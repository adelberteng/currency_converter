package db

import (
	redis "github.com/go-redis/redis/v8"

	"github.com/adelberteng/currency_converter/utils"
)


var cfg = utils.GetConfig()

func GetRedisClient() *redis.Client {
	redisEndpoint := cfg.Section("db").Key("redis_endpoint").String()
	redisPort := cfg.Section("db").Key("redis_port").String()

	r := redis.NewClient(&redis.Options{
		Addr:     redisEndpoint + ":" + redisPort,
		Password: "",
		DB:       0,
	})

	return r
}