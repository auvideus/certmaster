package main

import (
	"github.com/auvideus/certmaster"
	"github.com/auvideus/certmaster/cmd"
	log "github.com/sirupsen/logrus"
)

// Target of the certbot cleanup hook.
func main() {
	cmd.Initialize()

	config, err := certmaster.ReadYamlFile(certmaster.ConfigPath)
	if err != nil {
		log.Fatalln("could not read config file:", err)
	}

	err = certmaster.DeleteChallengeRecord(config)
	if err != nil {
		log.Fatalln("error deleting challenge record:", err)
	}
}
