package certmaster

import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"errors"
)

// PullCerts calls the rsync command internally to pull the certs.
func PullCerts(host string, path string) error {
	if host == "" || path == "" {
		return errors.New("host and path must be given")
	}
	var arguments []string
	arguments = append(arguments, "--archive")
	arguments = append(arguments, "--verbose")
	arguments = append(arguments, host + ":" + path + "/")
	arguments = append(arguments, "/etc/letsencrypt")

	log.Infoln("rsync arguments:", arguments)

	out, err := execCommand("rsync", arguments...).CombinedOutput()
	log.Infoln(string(out))
	if err != nil {
		return fmt.Errorf("error calling rsync command: %v", err)
	}
	return nil
}
