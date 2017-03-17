package levels

import (
	"os/exec"
	"strings"
)

type LevelOne struct {
}

func (l LevelOne) Check() (result bool, output string) {
	out, err := exec.Command("git", "rev-parse", "--git-dir").CombinedOutput()

	message := string(out)

	if err != nil || strings.Contains(message, "Not a git repository") {
        return false, message
    }

    return true, "You completed level 1!"
}

