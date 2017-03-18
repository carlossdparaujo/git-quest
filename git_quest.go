package main

import (
	"fmt"
	"git-quest/quest"
	"git-quest/quest/executors"
)

func main() {
	fmt.Printf("You are on Git Quest!\n")

	quest := quest.New("You've completed the quest!", "git", 
		[]string{"rev-parse", "--git-dir"}, executors.NewCommandLineExecutor())

	_, message := quest.Check()
	fmt.Printf("%s\n", message)
}