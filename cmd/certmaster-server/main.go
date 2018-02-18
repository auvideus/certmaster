package main

import (
	"os"
	"os/signal"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/auvideus/certmaster"
)

// Entrypoint of certmaster.  Calls certbot in a loop after reading the
// configuration file.
func main() {
	certmaster.Initialize()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		os.Exit(0)
	}()

	config, err := certmaster.ReadYamlFile(certmaster.ConfigPath)
	if err != nil {
		log.Fatalln("could not read config file", err)
	}

	duration, isZero := certmaster.GetServerPollInterval(config)
	log.Infoln("certmaster initialized")
	if !isZero {
		time.Sleep(10 * time.Second)
	}
	for {
		ok := certmaster.RefreshCerts(config)
		if !ok {
			log.Error("error refreshing certs")
		}
		if isZero {
			return
		}
		time.Sleep(duration)
	}
}
