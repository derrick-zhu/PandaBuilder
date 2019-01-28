package command

import (
	"PandaBuilder/logger"
	"PandaBuilder/models"
	"PandaBuilder/shell"
	"os"
	"path/filepath"

	"github.com/bclicn/color"
)

type CommandBootstrap struct {
	solutionData  *models.PandaSolutionModel
	solutionPath  string
	pandaFile     string
	pandaLockFile string
	result        bool
}

func NewCommandBootstrap(argv []string) *CommandBootstrap {
	result := &CommandBootstrap{}
	result.Start(argv)
	return result
}

func (c *CommandBootstrap) Start(param []string) bool {
	if len(param) < 3 {
		logger.Fatal("\n** Error: invalid parameters %v", param)
		return false
	}

	c.pandaFile = param[1]
	c.pandaLockFile = param[2]
	c.solutionData = &models.PandaSolutionModel{}
	var err error

	if filepath.IsAbs(c.pandaFile) == false {
		if c.pandaFile, err = filepath.Abs(c.pandaFile); err != nil {
			logger.Error("\n Error: could not get absolute directory path for %s", c.pandaFile)
			return false
		}
	}

	if filepath.IsAbs(c.pandaLockFile) == false {
		if c.pandaLockFile, err = filepath.Abs(c.pandaLockFile); err != nil {
			logger.Error("\n Error: could not get absolute directory path for %s", c.pandaLockFile)
			return false
		}
	}

	return true
}

func (c *CommandBootstrap) Run() {
	if _, err := os.Stat(c.pandaLockFile); os.IsNotExist(err) {
		logger.Println("Error: Could not find panda lock file. Please run \"%s\" first.", color.Red("PandaBuilder setup"))
		c.result = false
		return
	}

	if err := c.solutionData.LoadFrom(c.pandaFile); err != nil {
		logger.Fatal("\n** Error: fails in reading Pandfile: %s", color.Yellow(c.pandaFile))
		c.result = false
		return
	}

	if ret := c.solutionData.LoadFromLock(c.pandaLockFile); ret != models.Success {
		logger.Fatal("\n** Error: fails in reading Pandfile.lock: %s.", c.pandaLockFile)
		c.result = false
		return
	}

	for _, eachLock := range c.solutionData.Locks {
		pMod := c.solutionData.ModuleWithUrl(eachLock.URL())
		if pMod == nil {
			logger.Fatal("\n** Error: could not get module with %s.", eachLock.URL())
			c.result = false
			return
		}

		env := os.Environ()
		var workspace string
		var err error

		if workspace, err = filepath.Abs(pMod.Path); err != nil {
			logger.Println("warning: invalid module path %s.", workspace)
			continue
		}

		cmdStr := "git checkout " + eachLock.CommitHash
		var cmd = shell.ShellCommandWithEnv(cmdStr, workspace, env)
		if err = cmd.Run(); err != nil {
			logger.Fatal("\n** Error: fails in run command: %v.", err)
			continue
		}

		continue
	}

	return
}

func (c *CommandBootstrap) Done() int {
	return CmdResultSuccess
}
