package main

import (
	"github.com/auvideus/certmaster"
	"log"
	"os"
	"flag"
)

func main() {
	file := flag.String(
		"file",
		"/etc/certmaster/certmaster.yml",
		"full path of the configuration file to use")
	flag.Parse()

	config, err := certmaster.ReadYamlFile(*file)
	if err != nil {
		log.Fatal("error: %v", err)
		os.Exit(1)
	}

	err = certmaster.DeleteChallengeRecord(config)
	if err != nil {
		log.Fatalln("error deleting challenge record", err)
	}
}
