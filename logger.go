package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var critical = color.New(color.FgRed).PrintfFunc()
var info = color.New(color.FgBlue).PrintfFunc()
var success = color.New(color.FgGreen).PrintfFunc()
var warning = color.New(color.FgYellow).PrintfFunc()

func logMessage(f func(format string, a ...interface{}), status, prefix string, format string, a ...interface{}) {
	f("[%s] [%s] %s ", status, time.Now().Format("2006-01-02T15:04:05.000"), prefix)
	f(format, a...)
	f("\n")
}

func Fatal(format string, a ...interface{}) {
	logMessage(critical, "CRIT", "", format, a...)
	os.Exit(1)
}

func Critical(format string, a ...interface{}) {
	logMessage(critical, "CRIT", "", format, a...)
}

func Info(format string, a ...interface{}) {
	logMessage(info, "INFO", "", format, a...)
}

func Success(format string, a ...interface{}) {
	logMessage(success, "INFO", "", format, a...)
}

func Warning(format string, a ...interface{}) {
	logMessage(warning, "WARN", "", format, a...)
}

func Log(format string, a ...interface{}) {
	fmt.Printf("[%s]", time.Now().Format("2006-01-02T15:04:05.000"))
	fmt.Printf(format, a...)
	fmt.Printf("\n")
}
