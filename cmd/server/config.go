package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/pkg/constants"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	ServerPort string
	LogLevel   string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func GetConfig() *Config {
	return &Config{
		ServerPort: os.Getenv(constants.EnvServerPort),
		LogLevel:   os.Getenv(constants.EnvLogLevel),
		DBUser:     os.Getenv(constants.EnvDBUser),
		DBPassword: os.Getenv(constants.EnvDBPassword),
		DBHost:     os.Getenv(constants.EnvDBHost),
		DBPort:     os.Getenv(constants.EnvDBPort),
		DBName:     os.Getenv(constants.EnvDBName),
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
