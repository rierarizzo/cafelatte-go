package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/sirupsen/logrus"
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
	Logger(config.LogLevel)

	// Debug or release
	gin.SetMode(gin.ReleaseMode)
}

func Logger(logLevel string) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(false)

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.SetLevel(level)
}
