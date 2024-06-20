package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New()

func InitLogger() {
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only logrus the warning severity or above.
	logger.SetLevel(logrus.DebugLevel)
}

func GetLogger() *logrus.Logger {
	return logger
}
