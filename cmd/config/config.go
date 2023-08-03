package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/cmd/logger"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"os"
)

type Config struct {
	ServerPort string
	DSN        string
	LogLevel   string
}

func GetConfig() *Config {
	return &Config{
		ServerPort: os.Getenv(constants.EnvServerPort),
		DSN:        os.Getenv(constants.EnvDSN),
		LogLevel:   os.Getenv(constants.EnvLogLevel),
	}
}

func LoadInitConfig(config *Config) {
	// Config logger
	logger.ConfigLogger(config.LogLevel)

	// Debug or release
	gin.SetMode(gin.ReleaseMode)
}
