package main

import (
	"fmt"
	"strings"
    "os/exec"
)

func main() {
	fmt.Printf("You are on Git Quest!\n")

	out, err := exec.Command("git", "rev-parse", "--git-dir").CombinedOutput()

	message := string(out)

	if err != nil || strings.Contains(message, "Not a git repository") {
        fmt.Printf("%s", message)
        return
    }

    fmt.Printf("You completed level 1!\n")
}