package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/auvideus/certmaster"
	"github.com/auvideus/certmaster/cmd"
)

// Target of the certbot auth-hook script.
func main() {
	cmd.Initialize()

	config, err := certmaster.ReadYamlFile(certmaster.ConfigPath)
	if err != nil {
		log.Fatalln("could not read config file:", err)
	}

	err = certmaster.CreateChallengeRecord(config)
	if err != nil {
		log.Fatalln("error creating challenge record:", err)
	}
}
