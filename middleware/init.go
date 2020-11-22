
package middleware

import (
	"github.com/sirupsen/logrus"

	"github.com/c3sr/config"
	"github.com/c3sr/logger"
)

var (
	log *logrus.Entry
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "tracer/zipkin")
	})
}
