package utils

import (
	"log"
	"os"

	goLogger "github.com/adelberteng/go_logger"
)

func GetLogger() goLogger.Logger {
	os.MkdirAll(LoggerConf.Dir, 0766)
	logFile, err := os.OpenFile(LoggerConf.Dir + "/" + LoggerConf.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil && err == os.ErrNotExist {
		os.Create(LoggerConf.Dir + "/" + LoggerConf.FileName)
	} else if err != nil {
		log.Fatalf("log file open error : %v", err)
	}

	return goLogger.CreateLogger(logFile, "debug")
}