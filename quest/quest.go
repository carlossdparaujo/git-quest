package quest

import (
	"git-quest/quest/commands"
)

type quest struct {
	description string
	completionMessage string
	command commands.Command
	completed bool
}

func New(description string, completionMessage string, command commands.Command, completed bool) quest {
	return quest{description, completionMessage, command, completed}
}

func (q quest) Description() string {
	return q.description
}

func (q quest) Completed() bool {
	return q.completed
}

func (q *quest) Check(command commands.Command) (bool, string) {
	if (!q.command.Equals(command)) {
		return false, "Inserted the wrong command"
	}

	output, err := q.command.Execute()

	if (err != nil) {
		return false, output
	}

	q.completed = true
	return true, q.completionMessage	
}

