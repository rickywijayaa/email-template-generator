package log

import (
	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func init() {
	logger.SetFormatter(&log.JSONFormatter{})
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}
