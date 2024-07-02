package latihan_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("username", "aant").
		Info("Hello aant")
	logger.WithField("username", "aant").
		WithField("name", "sultan").
		Info("hello Aant sultan")

	logger.Info("hello logging")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"username": "aant",
		"name":     "sultan",
	}).Info("Hello fields")
}
