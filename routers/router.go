package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	config "gopkg.in/ini.v1"

	"github.com/adelberteng/currency_converter/handlers"
)

func Init() {
	cfg, err := config.Load("conf/config.ini")
	if err != nil {
		os.Exit(1)
	}
	servicePort := cfg.Section("app").Key("service_port").String()

	router := gin.Default()

	router.GET("/rate/:currency_type", handlers.GetCurrencyRate)
	router.POST("/rate", handlers.CountCurrencyRate)
	
	router.Run(":"+servicePort)
}