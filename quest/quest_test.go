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

type mockedDifferentCommand struct {
}

func (m mockedDifferentCommand) Execute() (out string, err error) {
	return "", nil
}

func (m mockedDifferentCommand) Equals(other commands.Command) bool {
	return false
}

func TestFailsWhenGivenCommandIfIsNotTheCorrectOne(t *testing.T) {
	quest := quest{
		"",
		"You completed the quest!",
		mockedDifferentCommand{},
		false}

	command := mockedDifferentCommand{}
	result, returnedMessage := quest.Check(command)

	assert.False(t, result)
	assert.Equal(t, "Inserted the wrong command", returnedMessage)
}

func TestQuestIsCompletedIfCommandSuccessfulyExecuted(t *testing.T) {
	command := mockedCommand{"", nil}

	quest := quest{
		"",
		"",
		command,
		false}

	quest.Check(command)

	assert.True(t, quest.Completed())
}

func TestQuestIsNotCompletedIfCommandIsCorrectButFailed(t *testing.T) {
	command := mockedCommand{"", errors.New("error string")}

	quest := quest{
		"",
		"",
		command,
		false}

	quest.Check(command)

	assert.False(t, quest.Completed())
}

func TestReturnsTrueAndCompletionMessageOnSuccess(t *testing.T) {
	completionMessage := "You completed the quest!"
	command := mockedCommand{"", nil}

	quest := quest{
		"",
		completionMessage,
		command,
		false}

	result, returnedMessage := quest.Check(command)

	assert.True(t, result)
	assert.Equal(t, completionMessage, returnedMessage)
}

func TestReturnsFalseAndOutputMessageOnFailure(t *testing.T) {
	output := "Failure output :("
	command := mockedCommand{output, errors.New("error string")}

	quest := quest{
		"",
		"You completed the quest!",
		command,
		false}

	result, returnedMessage := quest.Check(command)

	assert.False(t, result)
	assert.Equal(t, output, returnedMessage)
}