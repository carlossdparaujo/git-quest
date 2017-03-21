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
	output, err := q.command.Execute()

	if (err != nil) {
		return false, output
	}

	if (!q.command.Equals(command)) {
		return true, "Inserted the wrong command"
	} else {
		q.completed = true
		return true, q.completionMessage
	}
}

