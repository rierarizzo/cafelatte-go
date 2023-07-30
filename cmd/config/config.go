package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/cmd/logger"
	"github.com/rierarizzo/cafelatte/internal/domain/constants"
	"os"
)

type Config struct {
	ServerPort string
	DSN        string
	GinMode    string
	LogLevel   string
}

func GetConfig() *Config {
	return &Config{
		ServerPort: os.Getenv(constants.EnvServerPort),
		DSN:        os.Getenv(constants.EnvDSN),
		GinMode:    os.Getenv(constants.EnvGinMode),
		LogLevel:   os.Getenv(constants.EnvLogLevel),
	}
}

func LoadInitConfig(config *Config) {
	// Config logger
	logger.ConfigLogger()

	// Debug or release
	gin.SetMode(config.GinMode)
}
