package main

import (
	"github.com/auvideus/certmaster"
	log "github.com/sirupsen/logrus"
	"flag"
)

// Target of the certbot cleanup hook.
func main() {
	certmaster.Initialize()

	config, err := certmaster.ReadYamlFile(configPath)
	if err != nil {
		log.Fatalln("could not read config file:", err)
	}

	err = certmaster.DeleteChallengeRecord(config)
	if err != nil {
		log.Fatalln("error deleting challenge record:", err)
	}
}
