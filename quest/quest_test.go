package quest

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "errors"
)

var quest Quest

type mockedExecutor struct {
	output string
	err error
}

func (m mockedExecutor) execute(command string, args []string) (out string, err error) {
	return m.output, m.err
}

func TestReturnsTrueAndCompletionMessageOnSuccess(t *testing.T) {
	completionMessage := "You completed the quest!"

	quest := Quest{
		completionMessage,
		"command", 
		[]string{"args"}, 
		mockedExecutor{"Success output", nil}}

	result, returnedMessage := quest.Check()

	assert.True(t, result)
	assert.Equal(t, completionMessage, returnedMessage)
}

func TestReturnsFalseAndOutputMessageOnFailure(t *testing.T) {
	output := "Failure output :("

	quest := Quest{
		"You completed the quest!",
		"command", 
		[]string{"args"}, 
		mockedExecutor{output, errors.New("error string")}}

	result, returnedMessage := quest.Check()

	assert.False(t, result)
	assert.Equal(t, output, returnedMessage)
}