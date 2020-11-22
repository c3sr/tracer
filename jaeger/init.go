package jaeger

import (
	"github.com/sirupsen/logrus"

	"github.com/c3sr/config"
	logger "github.com/c3sr/logger"
)

type loggerWrapper struct {
	*logrus.Entry
}

var (
	log *loggerWrapper
)

func (l *loggerWrapper) Error(s string) {
	l.Entry.Error(s)
}

func init() {
	config.AfterInit(func() {
		log = &loggerWrapper{logger.New().WithField("pkg", "tracer/jaeger")}
	})
}
