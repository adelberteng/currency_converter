package handlers

import (
	"fmt"
	"math"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/adelberteng/currency_converter/models"
)

var ctx = context.Background()
var err error
var redis = models.GetRedisClient()


func GetCurrencyRate(c *gin.Context) {
	currency_type := c.Param("currency_type")
	target_type := c.Query("target_type")
	
	rateMap := models.GetRate(redis, currency_type)
	var val interface{}
	if target_type != "" {
		val = rateMap[target_type]
	} else {
		val = rateMap
	}

	c.JSON(http.StatusOK, gin.H{
		"currency_type": currency_type,
		"target_type":   target_type,
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

	currency_type := json["currency_type"]
	target_type := json["target_type"]
	amount_str := json["amount"]
	if currency_type == "" || target_type == "" {
		res.Message = "currency type and target_type is required."
		c.JSON(http.StatusBadRequest, gin.H{"message": res.Message})
		return 
	}

	amount, err := strconv.ParseInt(amount_str, 10, 64)
	if err != nil {
		fmt.Println(err)
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

	c.JSON(res.Status, gin.H{
		"message": res.Message,
		"result":  res.Result,
	})
}
