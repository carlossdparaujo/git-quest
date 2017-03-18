package quest

import (
	"git-quest/quest/executors"
)

type quest struct {
	completionMessage string
	command string
	args []string
	executor executors.Executor
}

func New(completionMessage string, command string, args []string, executor executors.Executor) quest {
	return quest{completionMessage, command, args, executor}
}

func (q quest) Check() (bool, string) {
	output, err := q.executor.Execute(q.command, q.args)

	if (err != nil) {
		return false, output
	}

	return true, q.completionMessage	
}

