package executors

import (
	"os/exec"
)

type CommandLineExecutor struct {
}

func (c CommandLineExecutor) Execute(command string, args []string) (string, error) {
	out, err := exec.Command(command, args...).CombinedOutput()
	return string(out), err
}