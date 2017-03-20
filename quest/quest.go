package quest

import (
	"git-quest/quest/commands"
)

type quest struct {
	completionMessage string
	command commands.Command
}

func New(completionMessage string, command commands.Command) quest {
	return quest{completionMessage, command}
}

func (q quest) Check() (bool, string) {
	output, err := q.command.Execute()

	if (err != nil) {
		return false, output
	}

	return true, q.completionMessage	
}

