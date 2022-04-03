package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mattn/go-isatty"
)

type Logger struct {
	prefix string
	file   *os.File
	color  bool
}

func New(prefix string, file *os.File) *Logger {
	var color bool

	if isatty.IsTerminal(file.Fd()) || isatty.IsCygwinTerminal(file.Fd()) {
		color = true
	}

	return &Logger{
		prefix: prefix,
		file:   file,
		color:  color,
	}
}

func (l *Logger) logMessage(f func(w io.Writer, format string, a ...interface{}), status, prefix string, format string, a ...interface{}) {
	if l.color {
		f(l.file, "[%s] [%s] %s ", status, time.Now().Format("2006-01-02T15:04:05.000"), prefix)
		f(l.file, format, a...)
		f(l.file, "\n")
	} else {
		fmt.Fprintf(l.file, "[%s] [%s] %s ", status, time.Now().Format("2006-01-02T15:04:05.000"), prefix)
		fmt.Fprintf(l.file, format, a...)
		fmt.Fprintf(l.file, "\n")
	}
}

func (l *Logger) Fatal(format string, a ...interface{}) {
	l.logMessage(critical, "CRIT", l.prefix, format, a...)
	os.Exit(1)
}

func (l *Logger) Critical(format string, a ...interface{}) {
	l.logMessage(critical, "CRIT", l.prefix, format, a...)
}

func (l *Logger) Info(format string, a ...interface{}) {
	l.logMessage(info, "INFO", l.prefix, format, a...)
}

func (l *Logger) Success(format string, a ...interface{}) {
	l.logMessage(success, "INFO", l.prefix, format, a...)
}

func (l *Logger) Warning(format string, a ...interface{}) {
	l.logMessage(warning, "WARN", l.prefix, format, a...)
}

func (l *Logger) Log(format string, a ...interface{}) {
	fmt.Fprintf(l.file, "[%s] %s", time.Now().Format("2006-01-02T15:04:05.000"), l.prefix)
	fmt.Fprintf(l.file, format, a...)
	fmt.Fprintf(l.file, "\n")
}
