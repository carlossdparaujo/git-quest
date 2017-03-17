package main

import (
	"fmt"
	"git-quest/levels"
)

func main() {
	fmt.Printf("You are on Git Quest!\n")

	l := levels.LevelOne{}

	_, message := l.Check()

	fmt.Printf("%s\n", message);
}