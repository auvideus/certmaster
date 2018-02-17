package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/auvideus/certmaster"
	"flag"
)

// Target of the certbot auth-hook script.
func main() {
	certmaster.Initialize()

	file := flag.String(
		"file",
		"/etc/certmaster/certmaster.yml",
		"full path of the configuration file to use")
	flag.Parse()

	config, err := certmaster.ReadYamlFile(*file)
	if err != nil {
		log.Fatalln("could not read config file:", err)
	}

	err = certmaster.CreateChallengeRecord(config)
	if err != nil {
		log.Fatalln("error creating challenge record:", err)
	}
}
