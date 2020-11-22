package zipkin

import (
	"github.com/sirupsen/logrus"

	"github.com/c3sr/config"
	"github.com/c3sr/logger"
)

type loggerWrapper struct {
	*logrus.Entry
}

var (
	log *loggerWrapper
)

func (l *loggerWrapper) Log(keyvals ...interface{}) error {
	l.Entry.Info(keyvals...)
	return nil
}

func init() {
	config.AfterInit(func() {
		log = &loggerWrapper{
			logger.New().WithField("pkg", "tracer/zipkin"),
		}
	})
}
