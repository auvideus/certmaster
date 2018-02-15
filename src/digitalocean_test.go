package src

import (
	"os"
	"testing"
	"time"
	"flag"
)

var (
	enabled = flag.Bool(
		"api",
		false,
		"run api tests")
)

func api(t *testing.T) {
	if !*enabled{
		t.Skip()
	}
}

func createConfig(t *testing.T) *Config {
	api(t)
	domain := os.Getenv("TEST_DOMAIN")
	email := os.Getenv("TEST_EMAIL")
	token := os.Getenv("TEST_TOKEN")
	os.Setenv("CERTBOT_DOMAIN", domain)
	os.Setenv("CERTBOT_VALIDATION", "testval")
	config := Config{
		Meta: Meta_S{
			Email: email,
		},
		Digital_Ocean: Digital_Ocean_S{
			Token: token,
		},
		Domains: []Domain_S{
			{
				Name: domain,
			},
		},
	}
	return &config
}

func TestCreateAndDelete(t *testing.T) {
	api(t)
	err := CreateChallengeRecord(createConfig(t))
	if err != nil {
		t.Error(err)
	}
	err = DeleteChallengeRecord(createConfig(t))
	if err != nil {
		t.Error(err)
	}
}
