package utils

import (
	"fmt"
	"os"

	config "gopkg.in/ini.v1"
)

func GetConfig() *config.File {
	cfg, err := config.Load("conf/config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cfg
}