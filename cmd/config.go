package cmd

import (
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"os"
)

type Config struct {
	ServerPort string
	DSN        string
	GinMode    string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort: os.Getenv(constants.EnvServerPort),
		DSN:        os.Getenv(constants.EnvDSN),
		GinMode:    os.Getenv(constants.EnvGinMode),
	}
}
