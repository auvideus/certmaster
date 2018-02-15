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
		log.Fatal("could not read config file: %v", err)
		os.Exit(1)
	}

	duration := getPollInterval(config)
	for {
		certbot.CallCertbot(*file, config)
		time.Sleep(duration)
	}
}

func getPollInterval(config *certbot.Config) (time.Duration) {
	duration, err := time.ParseDuration(config.Meta.PollInterval)
	if err != nil {
		log.Println("misconfigured poll interval, setting to 5s")
		duration, _ = time.ParseDuration("5s")
		return duration
	}
	tooShort := int64(duration) * 1000000 < 5
	if tooShort {
		log.Println(
			"poll interval too short, setting to 5s")
		duration, _ = time.ParseDuration("5s")
	}
	return duration
}
