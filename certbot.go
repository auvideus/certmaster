package certmaster

import (
	"os/exec"
	"log"
	"fmt"
)

type cmdI interface {
	CombinedOutput() ([]byte, error)
}

func cmdBuilder(name string, arg ...string) cmdI {
	return exec.Command(name, arg...)
}

var execCommand = cmdBuilder

func CallCertbot(file string, config *Config) error {
	for _, domain := range config.Domains {
		var arguments []string
		arguments = append(arguments, "certonly")
		arguments = append(arguments, "--non-interactive")
		arguments = append(arguments, "--manual-public-ip-logging-ok")
		arguments = append(arguments, "--agree-tos")
		arguments = append(arguments, "--email=" + config.Meta.Email)
		arguments = append(arguments, "--manual")
		arguments = append(arguments, "--preferred-challenges=dns")
		arguments = append(arguments, "--manual-auth-hook")
		arguments = append(arguments, "/opt/certmaster/pre.sh --file " + file)
		arguments = append(arguments, "--manual-cleanup-hook")
		arguments = append(arguments, "/opt/certmaster/post.sh --file " + file)
		arguments = append(arguments, "-d")
		arguments = append(arguments, domain.Name)
		for _, subdomain := range domain.Subdomains {
			arguments = append(arguments, "-d")
			arguments = append(arguments, subdomain)
		}

		allDomains := "\n\t" + domain.Name
		for _, subdomain := range domain.Subdomains {
			allDomains += "\n\t" + subdomain
		}
		log.Println("calling certbot for domains:" + allDomains)

		log.Println("certbot command arguments:", arguments)

		out, err := execCommand("certbot", arguments...).CombinedOutput()
		log.Println(string(out))
		if err != nil {
			return fmt.Errorf("error calling certbot command: %v", err)
		}
	}
	return nil
}
