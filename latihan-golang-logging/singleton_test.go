package latihan_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleton(t *testing.T) {
	logrus.Info("hello world")
	logrus.Warn("hello world")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("hello world")
	logrus.Warn("hello world")

}
