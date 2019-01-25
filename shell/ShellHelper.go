package shell

import (
	"os"
	"os/exec"
	"strings"
)

func ShellCommandWith(cmd string) *exec.Cmd {
	if len(cmd) <= 0 {
		return nil
	}

	var result *exec.Cmd

	cmdComp := strings.Fields(cmd)
	numOfComp := len(cmdComp)
	if numOfComp > 0 {
		program := cmdComp[0]
		var params []string

		if numOfComp > 1 {
			params = make([]string, numOfComp-1)
			copy(params, cmdComp[1:])
		}

		result = exec.Command(program, params...)
		result.Stdout = os.Stdout
		result.Stdin = os.Stdin
		result.Stderr = os.Stderr
	}
	return result
}
