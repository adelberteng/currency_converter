package utils

import (
	"log"
	"os"

	goLogger "github.com/adelberteng/go_logger"
)

var Logger goLogger.Logger

func init() {
	logDir := LoggerConf.Dir
	logName := LoggerConf.FileName
	logLevel := LoggerConf.Level

	os.MkdirAll(logDir, 0766)
	logFile, err := os.OpenFile(logDir+"/"+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil && err == os.ErrNotExist {
		os.Create(logDir + "/" + logName)
	} else if err != nil {
		log.Fatalf("log file open error : %v", err)
	}

	Logger = goLogger.CreateLogger(logFile, logLevel)
}