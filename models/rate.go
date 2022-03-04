package models

import (
	"context"

	"github.com/adelberteng/currency_converter/db"
	"github.com/adelberteng/currency_converter/utils"
)

var (
	logger = utils.Logger
	ctx    = context.Background()
	redis  = db.GetRedisClient()
)

func GetRate(currencyType, targetType string) (string, error) {
	val, err := redis.HGet(ctx, currencyType, targetType).Result()

	return val, err
}

func GetAllRate(currencyType string) (map[string]string, error) {
	val, err := redis.HGetAll(ctx, currencyType).Result()

	return val, err
}
