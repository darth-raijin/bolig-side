package utility

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
	DEBUG
)

func Log(level LogLevel, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	_, fileName := filepath.Split(file)

	logMsg := fmt.Sprintf(format, args...)

	switch level {
	case INFO:
		color.Set(color.FgCyan)
		log.Printf("[INFO] %s:%d - %s\n", fileName, line, logMsg)
		color.Unset()
	case WARNING:
		color.Set(color.FgYellow)
		log.Printf("[WARNING] %s:%d - %s\n", fileName, line, logMsg)
		color.Unset()
	case ERROR:
		color.Set(color.FgRed)
		log.Printf("[ERROR] %s:%d - %s\n", fileName, line, logMsg)
		color.Unset()
	case DEBUG:
		color.Set(color.FgGreen)
		log.Printf("[DEBUG] %s:%d - %s\n", fileName, line, logMsg)
		color.Unset()
	default:
		color.Set(color.FgWhite)
		log.Printf("[UNKNOWN] %s:%d - %s\n", fileName, line, logMsg)
		color.Unset()
	}
}
