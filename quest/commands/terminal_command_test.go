package commands

import (
  	"testing"
  	"github.com/stretchr/testify/assert"
)

type mockedCommand struct {
}

func (m mockedCommand) Execute() (out string, err error) {
	return "", nil
}

func (m mockedCommand) Equals(other Command) bool {
	return true
}

func TestIsNotEqualWhenCommandIsNotTerminalCommand(t *testing.T) {
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte{}, nil
	}

	firstCommand := terminalCommand{"", []string{}, mockedRunner}
	secondCommand := mockedCommand{}

	assert.False(t, firstCommand.Equals(secondCommand))
}

func TestIsEqualWhenCommandAndArgumentsAreTheSame(t *testing.T) {
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte{}, nil
	}

	firstCommand := terminalCommand{"", []string{}, mockedRunner}
	secondCommand := terminalCommand{"", []string{}, nil}

	assert.True(t, firstCommand.Equals(secondCommand))
}

func TestIsNotEqualWhenCommandIsDifferent(t *testing.T) {
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte{}, nil
	}

	firstCommand := terminalCommand{"", []string{}, mockedRunner}
	secondCommand := terminalCommand{"other", []string{}, nil}

	assert.False(t, firstCommand.Equals(secondCommand))
}

func TestIsNotEqualWhenArgumentsAreDifferent(t *testing.T) {
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte{}, nil
	}

	firstCommand := terminalCommand{"", []string{"argument"}, mockedRunner}
	secondCommand := terminalCommand{"", []string{"otherArgument"}, nil}

	assert.False(t, firstCommand.Equals(secondCommand))
}

func TestReturnsStringOutput(t *testing.T) {
	output := "Output string!"
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte(output), nil
	}

	terminalCommand := terminalCommand{"", []string{}, mockedRunner}

	returnedOutput, err := terminalCommand.Execute()

	assert.Equal(t, output, returnedOutput)
	assert.Empty(t, err)
}

func TestReturnsStringOutputWithoutNewlinesAtTheEnd(t *testing.T) {
	output := "Output string!"
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte(output + "\n"), nil
	}

	terminalCommand := terminalCommand{"", []string{}, mockedRunner}

	returnedOutput, err := terminalCommand.Execute()

	assert.Equal(t, output, returnedOutput)
	assert.Empty(t, err)
}

func TestReturnsStringOutputMaintainingNewlinesInTheMiddle(t *testing.T) {
	output := "Output string!\nAnd this is output too!\nThis is the last one!"
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte(output + "\n"), nil
	}

	terminalCommand := terminalCommand{"", []string{}, mockedRunner}

	returnedOutput, err := terminalCommand.Execute()

	assert.Equal(t, output, returnedOutput)
	assert.Empty(t, err)
}

