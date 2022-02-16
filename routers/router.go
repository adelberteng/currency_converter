package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adelberteng/currency_converter/handlers"
	"github.com/adelberteng/currency_converter/utils"
)

var logger = utils.GetLogger()


func SetupRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/rate/:currency_type", func(c *gin.Context) {
		currencyType := c.Param("currency_type")
		targetType := c.Query("target_type")

		val, err := handlers.GetCurrencyRate(currencyType, targetType)
		if err != nil {
			logger.Error(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"currency_type": currencyType,
			"target_type":   targetType,
			"exchange_rate": val,
		})
	})

	router.POST("/rate", func(c *gin.Context) {
		var json map[string]string
		c.BindJSON(&json)
		logger.Info(json)

		res, err := handlers.CountCurrencyRate(json)
		if err != nil {
			logger.Error(err)
		}

		c.JSON(res.Status, gin.H{
			"message": res.Message,
			"result":  res.Result,
		})
	})

	return router
}


