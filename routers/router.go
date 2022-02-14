package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/adelberteng/currency_converter/handlers"
	"github.com/adelberteng/currency_converter/utils"
)

var cfg = utils.GetConfig()

func Init() {
	router := gin.Default()

	router.GET("/rate/:currency_type", handlers.GetCurrencyRate)
	router.POST("/rate", handlers.CountCurrencyRate)

	servicePort := cfg.Section("app").Key("service_port").String()
	router.Run(":" + servicePort)
}
