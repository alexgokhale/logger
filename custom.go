package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	prefix string
}

func New(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
	}
}

func (l *Logger) Fatal(format string, a ...interface{}) {
	logMessage(critical, "CRIT", l.prefix, format, a...)
	os.Exit(1)
}

func (l *Logger) Critical(format string, a ...interface{}) {
	logMessage(critical, "CRIT", l.prefix, format, a...)
}

func (l *Logger) Info(format string, a ...interface{}) {
	logMessage(info, "INFO", l.prefix, format, a...)
}

func (l *Logger) Success(format string, a ...interface{}) {
	logMessage(success, "INFO", l.prefix, format, a...)
}

func (l *Logger) Warning(format string, a ...interface{}) {
	logMessage(warning, "WARN", l.prefix, format, a...)
}

func (l *Logger) Log(format string, a ...interface{}) {
	fmt.Printf("[%s] %s", time.Now().Format("2006-01-02T15:04:05.000"), l.prefix)
	fmt.Printf(format, a...)
	fmt.Printf("\n")
}
