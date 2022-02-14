package main

import (
	"os"
	"os/exec"

	"github.com/adelberteng/currency_converter/utils"
	"github.com/adelberteng/currency_converter/routers"
)

var cfg = utils.GetConfig()

func main() {
	logger := utils.GetLogger()
	cmd, err := exec.Command("python3", "rate_crawler/main.py").Output()
	if err != nil {
		logger.Error(err)
		logger.Error(string(cmd))
		os.Exit(1)
	}

	router := routers.SetupRoute()
	
	servicePort := cfg.Section("app").Key("service_port").String()
	router.Run(":" + servicePort)
}
