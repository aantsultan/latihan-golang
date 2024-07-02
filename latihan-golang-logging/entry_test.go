package latihan_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithFields(logrus.Fields{
		"username": "aant",
		"name":     "sultan",
	})
	entry.Info("heelo aant sultan")
}
