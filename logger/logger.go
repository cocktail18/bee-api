package logger

import (
	"go.uber.org/zap"
	"sync"
)

var logger *zap.Logger
var once = sync.Once{}

func SetLogger(l *zap.Logger) {
	logger = l
}

func GetLogger() *zap.Logger {
	once.Do(func() {
		if logger == nil {
			var err error
			logger, err = zap.NewDevelopment(zap.AddCaller())
			if err != nil {
				panic(err)
			}
		}
	})
	return logger
}
