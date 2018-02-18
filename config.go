package certmaster

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"time"
	log "github.com/sirupsen/logrus"
)

const (
	ConfigPath = "/etc/certmaster/certmaster.yml"
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

type Client struct {
	Poll_Interval string
	Dry_Run       bool
	Priv_Key 	  string
	Path          string
	Host          string
}

type Config struct {
	Server        Server
	Client	      Client
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
	return c, nil
}

// GetServerPollInterval checks that the poll interval is set to a default.
func GetServerPollInterval(config *Config) (
	duration time.Duration, isZero bool) {
	return getPollInterval("server", config)
}

// GetClientPollInterval checks that the poll interval is set to a default.
func GetClientPollInterval(config *Config) (
	duration time.Duration, isZero bool) {
	return getPollInterval("client", config)
}

// getPollInterval gets a polling interval value out of the config, setting
// it to a reasonable default if not set.
func getPollInterval(section string, config *Config) (
		duration time.Duration, isZero bool) {
	var timeValue string
	switch section {
	case "server":
		timeValue = config.Server.Poll_Interval
	case "client":
		timeValue = config.Client.Poll_Interval
	}
	if timeValue == "" {
		log.Infoln("running once, since no poll interval was set")
		return 0, true
	}
	duration, err := time.ParseDuration(timeValue)
	if err != nil {
		log.Infoln("misconfigured poll interval, setting to 5m")
		duration, _ = time.ParseDuration("5m")
		return duration, false
	}
	tooShort := int64(duration) * 1000000 < 5
	if tooShort {
		log.Infoln("poll interval too short, setting to 5m")
		duration, _ = time.ParseDuration("5m")
	}
	return duration, false
}
