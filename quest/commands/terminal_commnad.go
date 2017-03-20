package commands

import (
	"os/exec"
	"strings"
)

type terminalCommand struct {
	command string
	args []string
	runCommand runCommand
}

func NewTerminalCommand(command string, args []string) terminalCommand {
	runOnCommandLine := func (command string, args []string) ([]byte, error) {
			return exec.Command(command, args...).CombinedOutput()
		}

	return terminalCommand{command, args, runOnCommandLine}
}

func (c terminalCommand) Execute() (string, error) {
	out, err := c.runCommand(c.command, c.args)
	message := strings.Trim(string(out), "\n")
	return message, err
}