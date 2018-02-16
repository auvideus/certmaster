package certmaster

import (
	"testing"
	"os/exec"
)

type mockCmd struct {
	*exec.Cmd
}

func (c *mockCmd) CombinedOutput() ([]byte, error) {
	return []byte(c.Path + " executed"), nil
}

func mockCmdBuilder(name string, arg ...string) cmdI {
	return &mockCmd{exec.Command(name, arg...)}
}

func TestCallCertbot(t *testing.T) {
	execCommand = mockCmdBuilder
	defer func() {
		execCommand = cmdBuilder
	}()
	config := Config{
		Meta: Meta {
			Email: "example@example.com",
		},
		Domains: []Domain{
			{
				Name: "example.com",
				Subdomains: []string{},
			},
			{
				Name: "example2.com",
				Subdomains: []string{
					"app1.example2.com",
				},
			},
		},
	}
	CallCertbot("/certmaster.yml", &config)
}
