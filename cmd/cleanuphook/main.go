package main

import (
	"github.com/auvideus/certmaster"
	log "github.com/sirupsen/logrus"
	"flag"
)

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

	err = certmaster.DeleteChallengeRecord(config)
	if err != nil {
		log.Fatalln("error deleting challenge record:", err)
	}
}
