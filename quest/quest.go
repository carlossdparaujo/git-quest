package quest

import (
	"git-quest/quest/commands"
)

type quest struct {
	description string
	completionMessage string
	command commands.Command
}

func New(description string, completionMessage string, command commands.Command) quest {
	return quest{description, completionMessage, command}
}

func (q quest) Description() string {
	return q.description
}

func (q quest) Check() (bool, string) {
	output, err := q.command.Execute()

	if (err != nil) {
		return false, output
	}

	return true, q.completionMessage	
}

