package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"log/slog"
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
	levels := map[string]int{"debug": -4, "info": 0, "warn": 4, "error": 8}

	var programLevel = new(slog.LevelVar)
	programLevel.Set(slog.Level(levels[config.LogLevel]))
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))
	slog.SetDefault(logger)

	// Debug or release
	gin.SetMode(config.GinMode)
}
