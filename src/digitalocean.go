package src

import (
	"os"
	"github.com/digitalocean/godo"
	"github.com/digitalocean/godo/context"
	"log"
	"golang.org/x/oauth2"
	"strings"
	"regexp"
	"strconv"
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

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func CreateChallengeRecord(config *Config) {

	tokenSource := &TokenSource{
		AccessToken: config.Digital_Ocean.Token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	domain := os.Getenv("CERTBOT_DOMAIN")
	log.Println("creating record for " + domain)
	record, _, err := client.Domains.CreateRecord(
		context.TODO(),
		domain,
		&godo.DomainRecordEditRequest{
			Type:     recordType,
			Name:     recordNamePrefix,
			Data:     os.Getenv("CERTBOT_VALIDATION"),
			Priority: 0,
			Port:     0,
			TTL:      1800,
			Weight:   0,
			Flags:    0,
			Tag:      recordTag,
		},
	)
	if err == nil {
		log.Println(authPrefix + string(record.ID))
	}
}

func DeleteChallengeRecord(config *Config) {
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
		log.Fatal("no valid output found from auth script")
		return
	}

	domain := os.Getenv("CERTBOT_DOMAIN")
	log.Println("deleting record for " + domain)
	_, err := client.Domains.DeleteRecord(
		context.TODO(),
		domain,
		authOutputID,
	)
	if err != nil {
		log.Fatal("problem deleting record: ", err)
	}
}
