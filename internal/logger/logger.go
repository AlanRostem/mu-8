package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
)

var logFile = os.Stdout
var logLevel = LevelInfo

func SetFile(file *os.File) {
	logFile = file
}

func logPrintf(levelLabel, levelColor, format string, args ...any) {
	ts := time.Now().Format(time.RFC3339)
	_, err := fmt.Fprintf(logFile, "%s %s[%s]%s\t", ts, levelColor, levelLabel, colorReset)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(logFile, format+"\n", args...)
	if err != nil {
		panic(err)
	}
}

func Debugf(format string, args ...any) {
	if logLevel > LevelDebug {
		return
	}
	logPrintf("DEBUG", colorGreen, format, args...)
}

func Infof(format string, args ...any) {
	if logLevel > LevelInfo {
		return
	}
	logPrintf("INFO", colorBlue, format, args...)
}

func Warnf(format string, args ...any) {
	if logLevel > LevelWarn {
		return
	}
	logPrintf("WARN", colorYellow, format, args...)
}

func Errorf(format string, args ...any) {
	if logLevel > LevelError {
		return
	}
	logPrintf("ERROR", colorRed, format, args...)
}
