package main

import (
	"log"
	"os"
	"github.com/auvideus/certmaster"
	"flag"
)

func main() {
	file := flag.String("--file",
		"/etc/certmaster/certmaster.yml",
		"Full path of the configuration file to use.")

	config, err := certmaster.ReadYamlFile(*file)
	if err != nil {
		log.Fatal("error: %v", err)
		os.Exit(1)
	}

	err = certmaster.CreateChallengeRecord(config)
	if err != nil {
		log.Fatalln("error creating challenge record", err)
	}
}
