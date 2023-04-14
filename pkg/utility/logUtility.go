package utility

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

type logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

var loggerInstance = &logger{
	infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
	warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
	errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
}

func Log(level LogLevel, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	switch level {
	case INFO:
		loggerInstance.infoLogger.Println(message)
	case WARNING:
		loggerInstance.warningLogger.Println(message)
	case ERROR:
		loggerInstance.errorLogger.Println(message)
	}
}

func Info(format string, v ...interface{}) {
	Log(INFO, format, v...)
}

func Warning(format string, v ...interface{}) {
	Log(WARNING, format, v...)
}

func Error(format string, v ...interface{}) {
	Log(ERROR, format, v...)
}
