package main

import (
	"github.com/adelberteng/currency_converter/routers"
	"github.com/adelberteng/currency_converter/utils"
)

var (
	appConf = utils.AppConf
)

func main() {
	router := routers.SetupRoute()

	router.Run(":" + appConf.Port)
}
