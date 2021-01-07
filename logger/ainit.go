package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func InitLog() {

	Log = logrus.New()

	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetReportCaller(true)

}
