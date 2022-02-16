package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/adelberteng/currency_converter/handlers"
	"github.com/adelberteng/currency_converter/utils"
)

var logger = utils.GetLogger()


func SetupRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/rate/:currency_type", handlers.GetCurrencyRate)
	router.POST("/rate", handlers.CountCurrencyRate)

	return router
}


