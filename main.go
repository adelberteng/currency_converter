package main

import (
	"github.com/adelberteng/currency_converter/utils"
	"github.com/adelberteng/currency_converter/routers"
)

var (
	appConf = utils.AppConf
	logger = utils.GetLogger()
)

func main() {
	router := routers.SetupRoute()
	
	router.Run(":" + appConf.Port)
}
