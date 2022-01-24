package main

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var err error

func main() {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	router := gin.Default()

	router.GET("/rate/:currency_type", func(c *gin.Context) {
		currency_type := c.Param("currency_type")
		target_type := c.Query("target_type")
		
		var val interface{}
		if target_type == "" {
			val, err = redis.HGetAll(ctx, currency_type).Result()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			val, err = redis.HGet(ctx, currency_type, target_type).Result()
			if err != nil {
				fmt.Println(err)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"currency_type": currency_type,
			"target_type":   target_type,
			"exchange_rate": val,
		})
	})

	router.POST("/rate", func(c *gin.Context) {
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

		val_str, err := redis.HGet(ctx, currency_type, target_type).Result()
		if err != nil {
			fmt.Println(err)
		} else if val_str == "" {
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
	})

	router.Run(":8080")
}
