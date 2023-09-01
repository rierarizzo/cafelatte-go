package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	Port       = "SERVER_PORT"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBName     = "DB_NAME"
	LogLevel   = "LOG_LEVEL"
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
		ServerPort: os.Getenv(Port),
		LogLevel:   os.Getenv(LogLevel),
		DBUser:     os.Getenv(DBUser),
		DBPassword: os.Getenv(DBPassword),
		DBHost:     os.Getenv(DBHost),
		DBPort:     os.Getenv(DBPort),
		DBName:     os.Getenv(DBName),
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
