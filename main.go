package main

import (
	"github.com/adelberteng/currency_converter/utils"
	"github.com/adelberteng/currency_converter/routers"
)

var (
	cfg = utils.GetConfig()
	logger = utils.GetLogger()
)

func main() {
	router := routers.SetupRoute()
	
	servicePort := cfg.Section("app").Key("service_port").String()
	router.Run(":" + servicePort)
}
