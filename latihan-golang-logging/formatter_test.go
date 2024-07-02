package latihan_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("hello logger")
	logger.Warn("hello logger")
	logger.Error("hello logger")
}
