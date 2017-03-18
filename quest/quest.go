package quest

import (
	"git-quest/quest/executors"
)

type Quest struct {
	completionMessage string
	command string
	args []string
	executor executors.Executor
}

func (q Quest) Check() (bool, string) {
	output, err := q.executor.Execute(q.command, q.args)

	if (err != nil) {
		return false, output
	}

	return true, q.completionMessage	
}

