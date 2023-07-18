package cmd

import (
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"os"
)

type Config struct {
	DSN     string
	GinMode string
}

func LoadConfig() *Config {
	return &Config{
		DSN:     os.Getenv(constants.EnvDSN),
		GinMode: os.Getenv(constants.EnvGinMode),
	}
}
