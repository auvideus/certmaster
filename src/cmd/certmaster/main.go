package main

import (
	"os"
	"os/signal"
	"time"
	"log"
	certbot "github.com/auvideus/certmaster/src"
	"flag"
)

func main() {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		os.Exit(0)
	}()

	file := flag.String("--file",
		"/etc/certmaster/certmaster.yml",
		"Full path of the configuration file to use.")

	config, err := certbot.ReadYamlFile(*file)
	if err != nil {
		log.Fatal("error: %v", err)
		os.Exit(1)
	}

	for {
		certbot.CallCertbot(*file, config)
		time.Sleep(5 * time.Second)
	}
}
