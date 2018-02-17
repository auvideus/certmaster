package certmaster

import (
	"os/exec"
)

// MockCmd mocks an exec.Cmd object by storing one as an anonymous field.
type mockCmd struct {
	*exec.Cmd
}

// CombinedOutput allows the MockCmd type to implement CmdI.
func (c *mockCmd) CombinedOutput() ([]byte, error) {
	return []byte(c.Path + " executed"), nil
}

// MockCmdBuilder mocks getting a Cmd object.
func mockCmdBuilder(name string, arg ...string) cmdI {
	return &mockCmd{exec.Command(name, arg...)}
}
