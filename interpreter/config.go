package interpreter

import "github.com/AlanRostem/mu-8/internal/logger"

type Config struct {
	ClockSpeedHz int
	LogLevel     int
	LogFile      string
}

func NewDefaultConfig() Config {
	return Config{
		ClockSpeedHz: 700,
		LogLevel:     logger.LevelInfo,
		LogFile:      "",
	}
}
