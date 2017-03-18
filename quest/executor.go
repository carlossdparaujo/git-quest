package quest

type Executor interface {
	execute(command string, args []string) (out string, err error)
}