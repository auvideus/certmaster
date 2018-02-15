package main

import (
	"os"
	"os/signal"
	"time"
	"log"
	"github.com/auvideus/certmaster"
	"flag"
)

func main() {
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
		log.Fatal("could not read config file", err)
	}

	duration := certmaster.GetPollInterval(config)
	for {
		certmaster.CallCertbot(*file, config)
		time.Sleep(duration)
	}
}
