package certmaster

import (
	log "github.com/sirupsen/logrus"
	"fmt"
	"strings"
)

// RefreshCerts is used to get the certs for all domains in the configuration.
func RefreshCerts(file string, config *Config) (ok bool) {
	ok = true
	for _, domain := range config.Domains {
		allDomains := append([]string{domain.Name}, domain.Subdomains...)
		_, err := callCertbot(file, config.Server.Email, allDomains)
		if err != nil {
			ok = false
			fmt.Errorf("error refreshing certs for domain %v: %v",
				domain, err)
		}
	}
	return ok
}

// callCertbot actually calls the certbot command for the given information.
func callCertbot(file string, email string, domains []string) (
	args string, err error) {
	if file == "" || email == "" || len(domains) < 1 {
		return "", fmt.Errorf("file and email must not be empty")
	}
	var arguments []string
	arguments = append(arguments, "certonly")
	arguments = append(arguments, "--non-interactive")
	arguments = append(arguments, "--manual-public-ip-logging-ok")
	arguments = append(arguments, "--agree-tos")
	arguments = append(arguments, "--email=" + email)
	arguments = append(arguments, "--manual")
	arguments = append(arguments, "--preferred-challenges=dns")
	arguments = append(arguments, "--manual-auth-hook")
	arguments = append(arguments, "/etc/certmaster/pre.sh \"--file " + file + "\"")
	arguments = append(arguments, "--manual-cleanup-hook")
	arguments = append(arguments, "/etc/certmaster/post.sh \"--file " + file + "\"")
	for _, domain := range domains {
		arguments = append(arguments, "-d " + domain)
	}

	log.Infoln("certbot arguments:", arguments)

	out, err := execCommand("certbot", arguments...).CombinedOutput()
	log.Infoln(string(out))
	if err != nil {
		return "", fmt.Errorf(
			"error calling certbot command: %v", err)
	}
	return strings.Join(arguments, " "), nil
}
