package certmaster

import (
	"testing"
)

// TestCallCertbot verifies that the configuration is set appropriately.
func TestCallCertbot(t *testing.T) {
	execCommand = mockCmdBuilder
	defer func() {
		execCommand = cmdBuilder
	}()
	email := "example@example.com"
	domains := []string{
		"example.com",
		"app1.example.com",
		"app2.example.com",
	}
	_, err := callCertbot(
		email,
		domains,
	)
	if err != nil {
		t.Error("valid call was an error")
	}
	_, err = callCertbot(
		"",
		domains,
	)
	if err == nil {
		t.Error("empty email should be an error")
	}
	_, err = callCertbot(
		email,
		[]string{},
	)
	if err == nil {
		t.Error("empty domains should be an error")
	}
}

// TestRefreshCerts tests that the configuration send is appropriate.
func TestRefreshCerts(t *testing.T) {
	execCommand = mockCmdBuilder
	defer func() {
		execCommand = cmdBuilder
	}()
	config := Config{
		Server: Server{
			Email: "example@example.com",
		},
		Domains: []Domain{
			{
				Name: "example.com",
				Subdomains: []string{
					"app1.example.com",
					"app2.example.com",
				},
			},
			{
				Name: "example2.com",
				Subdomains: []string{
					"app1.example2.com",
					"app2.example2.com",
				},
			},
		},
	}
	ok := RefreshCerts(&config)
	if !ok {
		t.Error("valid call was an error")
	}
	ok = RefreshCerts(&Config{})
	if !ok {
		t.Error("no domains should be ok")
	}
}
