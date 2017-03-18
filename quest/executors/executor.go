package executors

type Executor interface {
	Execute(command string, args []string) (out string, err error)
}