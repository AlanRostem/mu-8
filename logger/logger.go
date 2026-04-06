package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	colorReset = "\033[0m"
	// colorRed    = "\033[31m"
	colorGreen = "\033[32m"
	// colorYellow = "\033[33m"
	colorBlue = "\033[34m"
)

const (
	LevelDebug = iota
	LevelInfo
)

var logFile = os.Stdout
var logLevel = LevelDebug

func SetFile(file *os.File) {
	logFile = file
}

func logPrintf(levelLabel, levelColor, format string, args ...any) {
	if logLevel > LevelInfo {
		return
	}
	ts := time.Now().Format(time.RFC3339)
	fmt.Fprintf(logFile, "%s %s[%s]%s\t", ts, levelColor, levelLabel, colorReset)
	fmt.Fprintf(logFile, format+"\n", args...)
}

func Debugf(format string, args ...any) {
	logPrintf("DEBUG", colorGreen, format, args...)
}

func Infof(format string, args ...any) {
	logPrintf("INFO", colorBlue, format, args...)
}
