package log

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Log() {
	log.SetFormatter(&logrus.JSONFormatter{}) // Log JSON
	log.SetLevel(logrus.InfoLevel)

	log.Info("Server is starting...")
	log.WithFields(logrus.Fields{
		"event": "user_login",
		"user":  "alice",
	}).Info("User logged in")
}
