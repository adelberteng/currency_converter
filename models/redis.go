package models

import (
	"context"
	"os"

	redis "github.com/go-redis/redis/v8"
	config "gopkg.in/ini.v1"

	"github.com/adelberteng/currency_converter/pkg"
)

var logger = pkg.GetLogger()
var ctx = context.Background()

func GetRedisClient() *redis.Client {
	cfg, err := config.Load("conf/config.ini")
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	redisEndpoint := cfg.Section("db").Key("redis_endpoint").String()
	redisPort := cfg.Section("db").Key("redis_port").String()

	return redis.NewClient(&redis.Options{
		Addr:     redisEndpoint + ":" + redisPort,
		Password: "",
		DB:       0,
	})
}

func GetRate(r *redis.Client, currencyType string) map[string]string {
	val, err := r.HGetAll(ctx, currencyType).Result()
	if err != nil {
		logger.Error(err)
	}
	return val
}
