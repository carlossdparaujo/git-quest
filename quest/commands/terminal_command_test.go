package commands

import (
  	"testing"
  	"github.com/stretchr/testify/assert"
)

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

