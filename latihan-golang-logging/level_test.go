package latihan_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.Trace("this is trace")
	logger.Debug("this is bug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("this is trace")
	logger.Debug("this is bug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}
