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
		log.Fatal("could not read config file: %v", err)
	}

	duration := getPollInterval(config)
	for {
		certmaster.CallCertbot(*file, config)
		time.Sleep(duration)
	}
}

func getPollInterval(config *certmaster.Config) (time.Duration) {
	duration, err := time.ParseDuration(config.Meta.Poll_Interval)
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
