package tracer

import (
	"github.com/sirupsen/logrus"

	"github.com/c3sr/config"
	"github.com/c3sr/logger"
)

var (
	log *logrus.Entry = logger.New().WithField("pkg", "tracer")
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "tracer")
	})
}
