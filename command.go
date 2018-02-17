package certmaster

import (
	"os/exec"
)

// cmdI is used as an interface for unit testing.
type cmdI interface {
	CombinedOutput() ([]byte, error)
}

// cmdBuilder is used to get a Cmd object via exec.Command.
func cmdBuilder(name string, arg ...string) cmdI {
	return exec.Command(name, arg...)
}

var execCommand = cmdBuilder
