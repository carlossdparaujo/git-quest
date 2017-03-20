package commands

type runCommand func(command string, args []string) ([]byte, error)