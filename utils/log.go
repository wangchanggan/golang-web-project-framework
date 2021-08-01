package utils

import (
	log "github.com/sirupsen/logrus"
	"golang-web-project-framework/config"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.Level(config.Config.LogLevel))
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args)
}
