package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/pkg/constants/env"
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
		ServerPort: os.Getenv(env.ServerPort),
		LogLevel:   os.Getenv(env.LogLevel),
		DBUser:     os.Getenv(env.DBUser),
		DBPassword: os.Getenv(env.DBPassword),
		DBHost:     os.Getenv(env.DBHost),
		DBPort:     os.Getenv(env.DBPort),
		DBName:     os.Getenv(env.DBName),
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
