package shell

import (
	"PandaBuilder/logger"
	"os"
	"os/exec"
	"strings"
)

func ShellCommandWith(cmd string) *exec.Cmd {
	return ShellCommandWithEnv(cmd, "", []string{})
}

func ShellCommandWithEnv(cmd string, workspace string, env []string) *exec.Cmd {
	if len(cmd) <= 0 {
		return nil
	}

	var result *exec.Cmd

	cmdComp := strings.Fields(cmd)
	numOfComp := len(cmdComp)
	if numOfComp > 0 {
		program := cmdComp[0]
		var err error

		if program, err = exec.LookPath(program); err != nil {
			logger.Fatal("\n** Error: could not execute %s.", program)
			return nil
		}

		var params []string

		if numOfComp > 1 {
			params = make([]string, numOfComp-1)
			copy(params, cmdComp[1:])
		}

		result = exec.Command(program, params...)
		result.Stdout = os.Stdout
		result.Stdin = os.Stdin
		result.Stderr = os.Stderr

		if len(workspace) > 0 {
			result.Dir = workspace
		}

		if len(env) > 0 {
			result.Env = env
		}
	}
	return result
}
