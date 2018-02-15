package main

import (
	certbot "github.com/auvideus/certmaster"
	"log"
	"os"
	"flag"
)

func main() {
	file := flag.String("--file",
		"/etc/certmaster/certmaster.yml",
		"Full path of the configuration file to use.")

	config, err := certbot.ReadYamlFile(*file)
	if err != nil {
		log.Fatal("error: %v", err)
		os.Exit(1)
	}

	certbot.DeleteChallengeRecord(config)
}
