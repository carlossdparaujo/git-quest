package executors

import (
	"os/exec"
	"strings"
)

type commandLineExecutor struct {
	runCommand runCommand
}

func NewCommandLineExecutor() commandLineExecutor {
	runOnCommandLine := func (command string, args []string) ([]byte, error) {
			return exec.Command(command, args...).CombinedOutput()
		}

	return commandLineExecutor{runOnCommandLine}
}

func (c commandLineExecutor) Execute(command string, args []string) (string, error) {
	out, err := c.runCommand(command, args)
	message := strings.Trim(string(out), "\n")
	return message, err
}