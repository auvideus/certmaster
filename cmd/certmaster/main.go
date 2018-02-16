package main

import (
	"os"
	"os/signal"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/auvideus/certmaster"
	"flag"
)

func main() {
	certmaster.Initialize()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		os.Exit(0)
	}()

	file := flag.String(
		"file",
		"/etc/certmaster/certmaster.yml",
		"full path of the configuration file to use")
	flag.Parse()

	config, err := certmaster.ReadYamlFile(*file)
	if err != nil {
		log.Fatalln("could not read config file", err)
	}

	duration := certmaster.GetPollInterval(config)
	time.Sleep(10 * time.Second)
	for {
		ok := certmaster.RefreshCerts(*file, config)
		if !ok {
			log.Infoln("error refreshing certs")
		}
		time.Sleep(duration)
	}
}
