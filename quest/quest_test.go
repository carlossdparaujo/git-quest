package quest

import (
  	"testing"
  	"github.com/stretchr/testify/assert"
  	"errors"
  	"git-quest/quest/commands"
)

type mockedCommand struct {
	output string
	err error
}

func (m mockedCommand) Execute() (out string, err error) {
	return m.output, m.err
}

func (m mockedCommand) Equals(other commands.Command) bool {
	return true
}

func TestReturnsTrueAndCompletionMessageOnSuccess(t *testing.T) {
	completionMessage := "You completed the quest!"

	quest := quest{
		"",
		completionMessage,
		mockedCommand{"Success output", nil}}

	result, returnedMessage := quest.Check()

	assert.True(t, result)
	assert.Equal(t, completionMessage, returnedMessage)
}

func TestReturnsFalseAndOutputMessageOnFailure(t *testing.T) {
	output := "Failure output :("

	quest := quest{
		"",
		"You completed the quest!",
		mockedCommand{output, errors.New("error string")}}

	result, returnedMessage := quest.Check()

	assert.False(t, result)
	assert.Equal(t, output, returnedMessage)
}