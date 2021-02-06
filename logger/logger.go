package logger

import (
	"fmt"

	"github.com/ahsan/todo/environment"
)

var level string = environment.GetLogLevel()

var logLevels = map[string]int{
	"emerg":   0,
	"alert":   1,
	"crit":    2,
	"error":   3,
	"warning": 4,
	"notice":  5,
	"info":    6,
	"":        6,
	"debug":   7,
}

func Error(logLine string) bool {
	if logLevels[environment.GetLogLevel()] <= 3 {
		fmt.Println(logLine)
	}
	return false
}

func Warning(logLine string) bool {
	if logLevels[environment.GetLogLevel()] <= 4 {
		fmt.Println(logLine)
	}
	return true
}

func Info(logLine string) bool {
	if logLevels[environment.GetLogLevel()] <= 6 {
		fmt.Println(logLine)
	}
	return true
}

func Debug(logLine string) {
	if logLevels[environment.GetLogLevel()] <= 7 {
		fmt.Println(logLine)
	}
}
