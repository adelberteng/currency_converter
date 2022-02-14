package main

import (
	"os"
	"os/exec"

	"github.com/adelberteng/currency_converter/utils"
	"github.com/adelberteng/currency_converter/routers"
)

func main() {
	logger := utils.GetLogger()
	cmd, err := exec.Command("python3", "rate_crawler/main.py").Output()
	if err != nil {
		logger.Error(err)
		logger.Error(string(cmd))
		os.Exit(1)
	}

	routers.Init()
}
