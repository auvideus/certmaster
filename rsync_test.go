package certmaster

import (
	"testing"
)

// TestCallCertbot verifies that the configuration is set appropriately.
func TestPullCerts(t *testing.T) {
	execCommand = mockCmdBuilder
	defer func() {
		execCommand = cmdBuilder
	}()
	host := "example.com"
	path := "/etc/letsencrypt"
	err := PullCerts(
		host,
		path,
	)
	if err != nil {
		t.Error("valid call was an error")
	}
	host = ""
	path = "/etc/letsencrypt"
	err = PullCerts(
		host,
		path,
	)
	if err == nil {
		t.Error("missing host returned ok")
	}
	host = "example.com"
	path = ""
	err = PullCerts(
		host,
		path,
	)
	if err == nil {
		t.Error("missing path returned ok")
	}
}
