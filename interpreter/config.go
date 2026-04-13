package interpreter

type Config struct {
	ClockSpeedHz int
}

func NewDefaultConfig() Config {
	return Config{
		ClockSpeedHz: 700,
	}
}
