package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"git-quest/quest"
	"git-quest/quest/commands"
)

func main() {
	fmt.Printf("You are on Git Quest!\n")

	command := commands.NewTerminalCommand("git", []string{"init"})
	quest := quest.New("You need to create a git repository!", "You've completed the quest!", command, false)

	fmt.Printf("%s\n", quest.Description())

	text := ""

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		text = strings.Trim(text, "\n")

		if (text == "exit") {
			break
		}

		command := commands.NewTerminalCommandFromText(text)

		_, message := quest.Check(command)
		fmt.Printf("%s\n", message)

		if (quest.Completed()) {
			break
		}
	}
}