package commands

type Command interface {
	Execute() (out string, err error)
}