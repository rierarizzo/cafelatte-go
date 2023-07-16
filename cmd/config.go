package cmd

import (
	"os"
)

type Config struct {
	DSN     string
	GinMode string
}

func LoadConfig() *Config {
	return &Config{
		DSN:     os.Getenv("DSN"),
		GinMode: os.Getenv("GIN_MODE"),
	}
}
