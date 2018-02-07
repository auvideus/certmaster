package files

import (
	"os"
	"os/exec"
	"os/signal"
	"fmt"
	"time"
	"github.com/digitalocean/godo"
	"github.com/digitalocean/godo/context"
	"golang.org/x/oauth2"
	"log"
)

const (
	pat = "eb5707df05339c916a804fff03a69074e6376aabcfb9bb89f6c3b1f7db7cdedf"
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

func main() {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		os.Exit(0)
	}()

	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	mode := os.Args[1]
	if mode == "pre" {
		pre(client)
	} else if mode == "post" {
		post(client)
	} else if mode != "loop" {
		fmt.Println("ERROR: invalid first argument")
		os.Exit(1)
	} else {
		for {
			fmt.Println("Calling certbot command...")
			out, err := exec.Command("certbot",
				"certonly",
				"--non-interactive",
				"--manual-public-ip-logging-ok",
				"--agree-tos",
				"--email=eahpublic@protonmail.com",
				//"--dry-run",
				"--manual",
				"--preferred-challenges=dns",
				"--manual-auth-hook",
				"/opt/pre.sh",
				"--manual-cleanup-hook",
				"/opt/post.sh",
				"-d",
				"auvideus.io").CombinedOutput()
			fmt.Println("%q\n", string(out))
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
			time.Sleep(5 * time.Second)
		}
	}
}

func pre(client *godo.Client) {
	fmt.Println("Creating DNS record...")
	context := context.TODO()
	client.Domains.CreateRecord(context, os.Getenv("CERTBOT_DOMAIN"), &godo.DomainRecordEditRequest{
		Type:     "TXT",
		Name:     "_acme-challenge.",
		Data:     os.Getenv("CERTBOT_VALIDATION"),
		Priority: 0,
		Port:     0,
		TTL:      1800,
		Weight:   0,
		Flags:    0,
		Tag:      "",
	})
	//client.Droplets.List(context, &godo.ListOptions{Page: 0, PerPage: 10})
}

func post(client *godo.Client) {
	fmt.Println("Removing DNS record...")
	//context := context.TODO()
}
