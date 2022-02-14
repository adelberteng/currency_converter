package models

import (
	"context"
	"os"

	redis "github.com/go-redis/redis/v8"
	config "gopkg.in/ini.v1"

	"github.com/adelberteng/currency_converter/utils"
)

var (
	logger = utils.GetLogger()
	ctx = context.Background()
)

type RateData struct {
	redis *redis.Client
}

func GetRateDataModel() RateData {
	cfg, err := config.Load("conf/config.ini")
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	redisEndpoint := cfg.Section("db").Key("redis_endpoint").String()
	redisPort := cfg.Section("db").Key("redis_port").String()

	redis := redis.NewClient(&redis.Options{
		Addr:     redisEndpoint + ":" + redisPort,
		Password: "",
		DB:       0,
	})

	return RateData{redis}
}

func (r *RateData) GetRate(currencyType, targetType string) (string, error) {
	val, err := r.redis.HGet(ctx, currencyType, targetType).Result()

	return val, err
}

func (r *RateData) GetAllRate(currencyType string) (map[string]string, error) {
	val, err := r.redis.HGetAll(ctx, currencyType).Result()

	return val, err
}
