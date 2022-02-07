package pkg

import (
    "fmt"
    "os"
    "log"

    config "gopkg.in/ini.v1"

    goLogger "github.com/adelberteng/go_logger"
)

func GetLogger() goLogger.Logger {
    cfg, err := config.Load("conf/config.ini")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    logDir := cfg.Section("log").Key("log_dir").String()
    logName := cfg.Section("log").Key("log_file_name").String()

    os.MkdirAll(logDir, 0666)
    logFile, err := os.OpenFile(logDir+"/"+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil && err == os.ErrNotExist {
        os.Create(logDir + "/" + logName)
    } else if err != nil {
        log.Fatalf("log file open error : %v", err)
    }
    
    logger := goLogger.CreateLogger(logFile, "debug")

    return logger
}

