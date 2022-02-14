package utils

import (
	"log"
	"os"

	goLogger "github.com/adelberteng/go_logger"
)

var cfg = GetConfig()

func GetLogger() goLogger.Logger {
	logDir := cfg.Section("log").Key("log_dir").String()
	logName := cfg.Section("log").Key("log_file_name").String()

	os.MkdirAll(logDir, 0766)
	logFile, err := os.OpenFile(logDir+"/"+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil && err == os.ErrNotExist {
		os.Create(logDir + "/" + logName)
	} else if err != nil {
		log.Fatalf("log file open error : %v", err)
	}

	return goLogger.CreateLogger(logFile, "debug")
}