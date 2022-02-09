package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/adelberteng/currency_converter/models"
	"github.com/adelberteng/currency_converter/pkg"
)

var logger = pkg.GetLogger()
var ctx = context.Background()
var err error
var redis = models.GetRedisClient()

func GetCurrencyRate(c *gin.Context) {
	currencyType := c.Param("currency_type")
	targetType := c.Query("target_type")

	rateMap := models.GetRate(redis, currencyType)
	var val interface{}
	if targetType != "" {
		val = rateMap[targetType]
	} else {
		val = rateMap
	}
	logger.Info("currencyType: " + currencyType + " targetType: " + targetType + " val: " + fmt.Sprint(val))

	c.JSON(http.StatusOK, gin.H{
		"currency_type": currencyType,
		"target_type":   targetType,
		"exchange_rate": val,
	})
}

func CountCurrencyRate(c *gin.Context) {
	var res struct {
		Status  int     `json:"status"`
		Message string  `json:"message"`
		Result  float64 `json:"result"`
	}

	var json map[string]string
	c.BindJSON(&json)
	logger.Info(json)

	currency_type := json["currency_type"]
	target_type := json["target_type"]
	amount_str := json["amount"]
	if currency_type == "" || target_type == "" {
		res.Message = "currency type and target_type is required."
		logger.Info(res)
		c.JSON(http.StatusBadRequest, gin.H{"message": res.Message})
		return
	}

	amount, err := strconv.ParseInt(amount_str, 10, 64)
	if err != nil {
		logger.Error(err)
	}

	rateMap := models.GetRate(redis, currency_type)
	val_str := rateMap[target_type]
	if val_str == "" {
		res.Status = http.StatusBadRequest
		res.Message = "currency type and target_type are not currect."
		c.JSON(res.Status, gin.H{"message": res.Message})
		return
	}
	val, err := strconv.ParseFloat(val_str, 64)

	res.Result = math.Round((float64(amount)*val)*1000) / 1000
	res.Message = "exchange complete."
	res.Status = http.StatusOK
	logger.Info(res)

	c.JSON(res.Status, gin.H{
		"message": res.Message,
		"result":  res.Result,
	})
}
