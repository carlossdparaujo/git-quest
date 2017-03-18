package executors

import (
  	"testing"
  	"github.com/stretchr/testify/assert"
)

func TestReturnsStringOutput(t *testing.T) {
	output := "Output string!"
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte(output), nil
	}

	commandLineExecutor := commandLineExecutor{mockedRunner}

	returnedOutput, err := commandLineExecutor.Execute("", []string{})

	assert.Equal(t, output, returnedOutput)
	assert.Empty(t, err)
}

func TestReturnsStringOutputWithoutNewlinesAtTheEnd(t *testing.T) {
	output := "Output string!"
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte(output + "\n"), nil
	}

	commandLineExecutor := commandLineExecutor{mockedRunner}

	returnedOutput, err := commandLineExecutor.Execute("", []string{})

	assert.Equal(t, output, returnedOutput)
	assert.Empty(t, err)
}

func TestReturnsStringOutputMaintainingNewlinesInTheMiddle(t *testing.T) {
	output := "Output string!\nAnd this is output too!\nThis is the last one!"
	mockedRunner := func(command string, args []string) ([]byte, error) {
		return []byte(output + "\n"), nil
	}

	commandLineExecutor := commandLineExecutor{mockedRunner}

	returnedOutput, err := commandLineExecutor.Execute("", []string{})

	assert.Equal(t, output, returnedOutput)
	assert.Empty(t, err)
}

