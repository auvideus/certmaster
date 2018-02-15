package certmaster

import (
	"os/exec"
	"log"
)

func CallCertbot(file string, config *Config) {
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
		arguments = append(arguments, "/opt/pre.sh --file " + file)
		arguments = append(arguments, "--manual-cleanup-hook")
		arguments = append(arguments, "/opt/post.sh --file " + file)
		arguments = append(arguments, "-d")
		arguments = append(arguments, domain.Name)
		for _, subdomain := range domain.Subdomains {
			arguments = append(arguments, "-d")
			arguments = append(arguments, subdomain)
		}

		alldomains := "\t" + domain.Name
		for _, subdomain := range domain.Subdomains {
			alldomains += "\n\t" + subdomain
		}
		log.Println("calling certbot for domains\n" + alldomains)

		log.Println("certbot command arguments: %v", arguments)

		out, err := exec.Command("certbot", arguments...).CombinedOutput()
		log.Println("%q\n", string(out))
		if err != nil {
			log.Fatal(err)
		}
	}
}
