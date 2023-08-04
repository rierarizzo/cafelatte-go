package config

import (
	"github.com/sirupsen/logrus"
)

func ConfigLogger(logLevel string) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(false)

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.SetLevel(level)
}
