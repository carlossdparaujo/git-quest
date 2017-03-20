package commands

import (
	"os/exec"
	"strings"
	"git-quest/comparators"
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

func (c terminalCommand) Equals(other Command) bool {
	otherTerminalCommand, isTerminalCommand := other.(terminalCommand)

	if (!isTerminalCommand) {
		return false
	}

	return (c.command == otherTerminalCommand.command) && 
		comparators.CompareStringArrays(c.args, otherTerminalCommand.args)
}