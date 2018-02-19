package certmaster

import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"errors"
	"strings"
	"os/exec"
)

// PullCerts calls the rsync command internally to pull the certs.
func PullCerts(host string, path string, dryRun bool) error {
	if host == "" || path == "" {
		return errors.New("host and path must be given")
	}
	sshPath, _ := exec.LookPath("ssh")
	var arguments []string
	arguments = append(arguments, "--archive")
	arguments = append(arguments, "--verbose")
	arguments = append(arguments, "--itemize-changes")
	arguments = append(arguments, strings.Join(
		[]string{
			"-e \" " + sshPath,
			"-o StrictHostKeyChecking=no",
			"-o BatchMode=yes",
			"-o IdentityFile=/root/.ssh/id_rsa_certmaster",
			"-o UserKnownHostsFile=/dev/null",
			"\""},
		" "))
	arguments = append(arguments, host + ":" + path + "/")
	arguments = append(arguments, "/etc/letsencrypt")

	log.Infoln("rsync arguments:", arguments)

	if dryRun {
		return nil
	}
	out, err := execCommand("rsync", arguments...).CombinedOutput()
	log.Infoln(string(out))
	if err != nil {
		return fmt.Errorf("error calling rsync command: %v", err)
	}
	return nil
}
