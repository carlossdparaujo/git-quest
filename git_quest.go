package main

import (
	"fmt"
	"git-quest/quest"
	"git-quest/quest/commands"
)

func main() {
	fmt.Printf("You are on Git Quest!\n")

	command := commands.NewTerminalCommand("git", []string{"rev-parse", "--git-dir"})
	quest := quest.New("You've completed the quest!", command)

	_, message := quest.Check()
	fmt.Printf("%s\n", message)
}