package db

import (
	redis "github.com/go-redis/redis/v8"

	"github.com/adelberteng/currency_converter/utils"
)


var redisConf = utils.RedisConf

func GetRedisClient() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     redisConf.Endpoint + ":" + redisConf.Port,
		Password: "",
		DB:       0,
	})

	return r
}