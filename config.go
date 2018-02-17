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

type Server struct {
	Email string
	Poll_Interval string
	Dry_Run bool
}

type Config struct {
	Server        Server
	Digital_Ocean DigitalOcean
	Domains       []Domain
}

// ReadYamlFile tests that the file can be read and the configuration is
// correct.
func ReadYamlFile(file string) (c *Config, err error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err2 := yaml.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err2
	}
	if c.Server.Email == "" {
		return nil, errors.New(
			"email missing from configuration file")
	}
	if c.Server.Email == "" {
		return nil, errors.New(
			"token missing from configuration file")
	}
	if len(c.Domains) < 1 {
		return nil, errors.New(
			"no domains specified in configuration")
	}
	return c, nil
}

// GetPollInterval checks that the poll interval is set to a default.
func GetPollInterval(config *Config) (time.Duration) {
	duration, err := time.ParseDuration(config.Server.Poll_Interval)
	if err != nil {
		log.Infoln("misconfigured poll interval, setting to 5m")
		duration, _ = time.ParseDuration("5m")
		return duration
	}
	tooShort := int64(duration) * 1000000 < 5
	if tooShort {
		log.Infoln("poll interval too short, setting to 5m")
		duration, _ = time.ParseDuration("5m")
	}
	return duration
}
