package certmaster

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"time"
	"log"
)

type Domain_S struct {
	Name string
	Subdomains []string
}

type Digital_Ocean_S struct {
	Token string
}

type Meta_S struct {
	Email string
	Poll_Interval string
	Dry_Run bool
}

type Config struct {
	Meta Meta_S
	Digital_Ocean Digital_Ocean_S
	Domains []Domain_S
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
	return c, nil
}

func GetPollInterval(config *Config) (time.Duration) {
	duration, err := time.ParseDuration(config.Meta.Poll_Interval)
	if err != nil {
		log.Println("misconfigured poll interval, setting to 5s")
		duration, _ = time.ParseDuration("5s")
		return duration
	}
	tooShort := int64(duration) * 1000000 < 5
	if tooShort {
		log.Println("poll interval too short, setting to 5s")
		duration, _ = time.ParseDuration("5s")
	}
	return duration
}
