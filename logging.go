package certmaster

import (
	"gopkg.in/gemnasium/logrus-graylog-hook.v2"
	log "github.com/sirupsen/logrus"
	"os"
)

// Initialize will enable Graylog testing if the environment is set.
func Initialize() {
	graylogEndpoint := os.Getenv("GRAYLOG_ENDPOINT")
	if graylogEndpoint == "" {
		return
	}
	hook := graylog.NewGraylogHook(graylogEndpoint, nil)
	log.AddHook(hook)
}
