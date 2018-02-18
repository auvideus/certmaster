package cmd

import (
	log "github.com/sirupsen/logrus"
	"os"
	"gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

// Initialize will enable Graylog testing if the environment is set.
func Initialize() {
	log.SetOutput(os.Stdout)
	graylogEndpoint := os.Getenv("GRAYLOG_ENDPOINT")
	if graylogEndpoint == "" {
		return
	}
	hook := graylog.NewGraylogHook(graylogEndpoint, nil)
	log.AddHook(hook)
}
