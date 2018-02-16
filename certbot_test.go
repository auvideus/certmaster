package certmaster

import (
	"testing"
	"os/exec"
)

// mockCmd mocks an exec.Cmd object by storing one as an anonymous field.
type mockCmd struct {
	*exec.Cmd
}

// CombinedOutput allows the mockCmd type to implement CmdI.
func (c *mockCmd) CombinedOutput() ([]byte, error) {
	return []byte(c.Path + " executed"), nil
}

// mockCmdBuilder mocks getting a Cmd object.
func mockCmdBuilder(name string, arg ...string) cmdI {
	return &mockCmd{exec.Command(name, arg...)}
}

// TestCallCertbot verifies that the configuration is set appropriately.
func TestCallCertbot(t *testing.T) {
	execCommand = mockCmdBuilder
	defer func() {
		execCommand = cmdBuilder
	}()
	file := "/certmaster.yml"
	email := "example@example.com"
	domains := []string{
		"example.com",
		"app1.example.com",
		"app2.example.com",
	}
	_, err := callCertbot(
		file,
		email,
		domains,
	)
	if err != nil {
		t.Error("valid call was an error")
	}
	_, err = callCertbot(
		"",
		email,
		domains,
	)
	if err == nil {
		t.Error("empty file should be an error")
	}
	_, err = callCertbot(
		file,
		"",
		domains,
	)
	if err == nil {
		t.Error("empty email should be an error")
	}
	_, err = callCertbot(
		file,
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
		Meta: Meta{
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
	ok := RefreshCerts("/test.yml", &config)
	if !ok {
		t.Error("valid call was an error")
	}
	ok = RefreshCerts("", &config)
	if ok {
		t.Error("empty file should be an error")
	}
	ok = RefreshCerts("/test.yml", &Config{})
	if !ok {
		t.Error("no domains should be ok")
	}
}
