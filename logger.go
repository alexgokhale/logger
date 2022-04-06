package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

var critical = color.New(color.FgRed).FprintfFunc()
var info = color.New(color.FgBlue).FprintfFunc()
var success = color.New(color.FgGreen).FprintfFunc()
var warning = color.New(color.FgYellow).FprintfFunc()

func logMessage(f func(w io.Writer, format string, a ...interface{}), status string, format string, a ...interface{}) {
	f(os.Stdout, "[%s] [%s] ", status, time.Now().Format("2006-01-02T15:04:05.000"))
	f(os.Stdout, format, a...)
	f(os.Stdout, "\n")
}

func Fatal(format string, a ...interface{}) {
	logMessage(critical, "CRIT", format, a...)
	os.Exit(1)
}

func Critical(format string, a ...interface{}) {
	logMessage(critical, "CRIT", format, a...)
}

func Info(format string, a ...interface{}) {
	logMessage(info, "INFO", format, a...)
}

func Success(format string, a ...interface{}) {
	logMessage(success, "INFO", format, a...)
}

func Warning(format string, a ...interface{}) {
	logMessage(warning, "WARN", format, a...)
}

func Log(format string, a ...interface{}) {
	fmt.Printf("[%s]", time.Now().Format("2006-01-02T15:04:05.000"))
	fmt.Printf(format, a...)
	fmt.Printf("\n")
}
