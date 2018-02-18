package main

import (
	"os"
	"os/signal"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/auvideus/certmaster"
	"github.com/auvideus/certmaster/cmd"
)

// Entrypoint of certmaster.  Calls certbot in a loop after reading the
// configuration file.
func main() {
	cmd.Initialize()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		os.Exit(0)
	}()

	config, err := certmaster.ReadYamlFile(certmaster.ConfigPath)
	if err != nil {
		log.Fatalln("could not read config file:", err)
	}

	duration, isZero := certmaster.GetClientPollInterval(config)
	log.Infoln("certmaster initialized")
	if !isZero {
		time.Sleep(10 * time.Second)
	}
	for {
		err := certmaster.PullCerts(
			config.Client.Host, config.Client.Path, config.Client.Dry_Run)
		if err != nil {
			log.Error("error pulling certs:", err)
		}
		if isZero {
			return
		}
		time.Sleep(duration)
	}
}
