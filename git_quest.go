package main

import (
	"fmt"
	"git-quest/level"
)

func main() {
	fmt.Printf("You are on Git Quest!\n")

	l := levels.Level{Command: "git", Args: []string{"rev-parse", "--git-dir"}}

	_, message := l.Check()

	fmt.Printf("%s\n", message);
}