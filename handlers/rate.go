package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/adelberteng/currency_converter/models"
	"github.com/adelberteng/currency_converter/utils"
	"github.com/gin-gonic/gin"
)

var (
	logger = utils.GetLogger()
)

func GetCurrencyRate(c *gin.Context) {
	currencyType := c.Param("currency_type")
	targetType := c.Query("target_type")

	var val interface{}
	var err error

	if targetType != "" {
		val, err = models.GetRate(currencyType, targetType)
	} else {
		val, err = models.GetAllRate(currencyType)
	}
	if err != nil {
		logger.Error(err)
	}

	logger.Info("currencyType: " + currencyType + " targetType: " + targetType + " val: " + fmt.Sprint(val))

	c.JSON(http.StatusOK, gin.H{
		"currency_type": currencyType,
		"target_type":   targetType,
		"exchange_rate": val,
	})
}

func CountCurrencyRate(c *gin.Context) {
	var json map[string]string
	c.ShouldBindJSON(&json)
	logger.Info(json)

	currencyType := json["currency_type"]
	targetType := json["target_type"]
	amountStr := json["amount"]
	if currencyType == "" || targetType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "currency type and target_type is required.",
			"result":  "",
		})
		return
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"result":  "",
		})
		logger.Error(err)
		return
	}

	rateStr, err := models.GetRate(currencyType, targetType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"result":  "",
		})
		logger.Error(err)
		return 
	}

	if rateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "currency type and target_type are not currect.",
			"result":  "",
		})
		return
	}
	val, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"result":  "",
		})
		logger.Error(err)
		return
	}

	newAmount := math.Round((float64(amount)*val)*1000) / 1000

	c.JSON(http.StatusOK, gin.H{
		"message": "exchange complete.",
		"result":  newAmount,
	})
}
