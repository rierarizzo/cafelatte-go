package server

import (
	"os"

	"github.com/rierarizzo/cafelatte/pkg/constants/env"
	"github.com/sirupsen/logrus"
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
	Logger(config.LogLevel)
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
