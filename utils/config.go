package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type RedisConfig struct {
	Endpoint string
	Port string
}

type AppConfig struct {
	Port string
}

type LoggerConfig struct {
	Dir string
	FileName string
}

var (
	RedisConf RedisConfig
	AppConf AppConfig
	LoggerConf LoggerConfig
)


func init() {
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	RedisConf.Endpoint = cfg.Section("db").Key("redis_endpoint").String()
	RedisConf.Port = cfg.Section("db").Key("redis_port").String()

	AppConf.Port = cfg.Section("app").Key("service_port").String()

	LoggerConf.Dir = cfg.Section("log").Key("log_dir").String()
	LoggerConf.FileName = cfg.Section("log").Key("log_file_name").String()
}
