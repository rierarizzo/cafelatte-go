package logger

import (
	"github.com/sirupsen/logrus"
)

func ConfigLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
}
