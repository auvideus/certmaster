package certmaster

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"time"
	log "github.com/sirupsen/logrus"
	"errors"
)

type Domain struct {
	Name string
	Subdomains []string
}

type DigitalOcean struct {
	Token string
}

type Meta struct {
	Email string
	Poll_Interval string
	Dry_Run bool
}

type Config struct {
	Meta          Meta
	Digital_Ocean DigitalOcean
	Domains       []Domain
}

func ReadYamlFile(file string) (c *Config, err error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err2 := yaml.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err2
	}
	if c.Meta.Email == "" {
		return nil, errors.New(
			"email missing from configuration file")
	}
	if c.Meta.Email == "" {
		return nil, errors.New(
			"token missing from configuration file")
	}
	if len(c.Domains) < 1 {
		return nil, errors.New(
			"no domains specified in configuration")
	}
	return c, nil
}

func GetPollInterval(config *Config) (time.Duration) {
	duration, err := time.ParseDuration(config.Meta.Poll_Interval)
	if err != nil {
		log.Infoln("misconfigured poll interval, setting to 5s")
		duration, _ = time.ParseDuration("5s")
		return duration
	}
	tooShort := int64(duration) * 1000000 < 5
	if tooShort {
		log.Infoln("poll interval too short, setting to 5s")
		duration, _ = time.ParseDuration("5s")
	}
	return duration
}
