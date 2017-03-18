package quest

import (
	"git-quest/quest/executors"
)

type Quest struct {
	CompletionMessage string
	Command string
	Args []string
	Executor executors.Executor
}

func (q Quest) Check() (bool, string) {
	output, err := q.Executor.Execute(q.Command, q.Args)

	if (err != nil) {
		return false, output
	}

	return true, q.CompletionMessage	
}

