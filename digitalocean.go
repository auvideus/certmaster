package certmaster

import (
	"os"
	"github.com/digitalocean/godo"
	"github.com/digitalocean/godo/context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"strings"
	"regexp"
	"strconv"
	"errors"
)

const (
	authPrefix = "authhook_result: "
	authPattern = ".*" + authPrefix + "(.*)"
	recordTag = "certmaster"
	recordNamePrefix = "_acme-challenge"
	recordType = "TXT"
)

type TokenSource struct {
	AccessToken string
}

// Token creates an oauth2 token.
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// CreateChallengeRecord creates a Digital Ocean challenge record.
func CreateChallengeRecord(config *Config) error {

	tokenSource := &TokenSource{
		AccessToken: config.Digital_Ocean.Token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	domain := os.Getenv("CERTBOT_DOMAIN")
	data := os.Getenv("CERTBOT_VALIDATION")
	if domain == "" {
		return errors.New(
			"missing CERTBOT_DOMAIN environment variable")
	}
	if data == "" {
		return errors.New(
			"missing CERTBOT_VALIDATION environment variable")
	}
	log.Infoln("creating record for " + domain)

	if config.Server.Dry_Run {
		log.Infoln("because of dry run, not actually creating record")
		log.Infoln(authPrefix + "00000000")
		return nil
	}
	record, _, err := client.Domains.CreateRecord(
		context.TODO(),
		domain,
		&godo.DomainRecordEditRequest{
			Type:     recordType,
			Name:     recordNamePrefix,
			Data:     data,
			Priority: 0,
			Port:     0,
			TTL:      1800,
			Weight:   0,
			Flags:    0,
			Tag:      recordTag,
		},
	)
	if err == nil {
		// for unit testing
		os.Setenv(
			"CERTBOT_AUTH_OUTPUT", authPrefix + strconv.Itoa(record.ID))

		log.Infoln(authPrefix + strconv.Itoa(record.ID))
		return nil
	}
	return err
}

// DeleteChallengeRecord deletes the previously-created challenge record that
// was part of the same certbot invocation.
func DeleteChallengeRecord(config *Config) error {
	tokenSource := &TokenSource{
		AccessToken: config.Digital_Ocean.Token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	authOutputID := -1
	authOutput := os.Getenv("CERTBOT_AUTH_OUTPUT")
	authOutputLines := strings.Split(authOutput, "\n")
	re := regexp.MustCompile(authPattern)
	for _, line := range authOutputLines {
		if re.MatchString(line) {
			authOutputID, _ = strconv.Atoi(re.FindStringSubmatch(line)[1])
			break
		}
	}

	if authOutputID == -1 {
		return errors.New("no valid output found from auth script")
	}

	domain := os.Getenv("CERTBOT_DOMAIN")
	log.Infoln("deleting record for " + domain)
	if config.Server.Dry_Run {
		log.Infoln("because of dry run, not actually deleting record")
		return nil
	}
	_, err := client.Domains.DeleteRecord(
		context.TODO(),
		domain,
		authOutputID,
	)
	return err
}
