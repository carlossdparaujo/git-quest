package levels

import (
	"os/exec"
)

type Level struct {
	Command string
	Args []string
}

func (l Level) Check() (result bool, output string) {
	out, err := exec.Command(l.Command, l.Args...).CombinedOutput()

	message := string(out)

	if err != nil {
        return false, message
    }

    return true, "You completed level 1!"
}

