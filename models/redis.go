package models

import (
	"os"

	"github.com/go-redis/redis/v8"
	config "gopkg.in/ini.v1"
)


func GetRedisClient() *redis.Client {
	cfg, err := config.Load("conf/config.ini")
	if err != nil {
		os.Exit(1)
	}
	redisEndpoint := cfg.Section("db").Key("redis_endpoint").String()
	redisPort := cfg.Section("db").Key("redis_port").String()
	r := redis.NewClient(&redis.Options{
		Addr:     redisEndpoint+":"+redisPort,
		Password: "",
		DB:       0,
	})

	return r
}