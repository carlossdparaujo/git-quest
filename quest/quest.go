package quest

type Quest struct {
	completionMessage string
	command string
	args []string
	executor Executor
}

func (q Quest) Check() (bool, string) {
	output, err := q.executor.execute(q.command, q.args)

	if (err != nil) {
		return false, output
	}

	return true, q.completionMessage	
}

