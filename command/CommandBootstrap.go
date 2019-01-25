package command

import (
	"PandaBuilder/models"
	"fmt"
	"log"
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
		log.Fatalf("\n** error: invalid parameters %v", param)
		return false
	}

	c.pandaFile = param[1]
	c.pandaLockFile = param[2]
	c.solutionData = &models.PandaSolutionModel{}
	var err error

	if filepath.IsAbs(c.pandaFile) == false {
		if c.pandaFile, err = filepath.Abs(c.pandaFile); err != nil {
			log.Printf("\n** error: could not get absolute directory path for %s", c.pandaFile)
			return false
		}
	}

	if filepath.IsAbs(c.pandaLockFile) == false {
		if c.pandaLockFile, err = filepath.Abs(c.pandaLockFile); err != nil {
			log.Printf("\n** error: could not get absolute directory path for %s", c.pandaLockFile)
			return false
		}
	}

	return true
}

func (c *CommandBootstrap) Run() {
	if _, err := os.Stat(c.pandaLockFile); os.IsNotExist(err) {
		fmt.Printf("\n** Error: Could not find panda lock file. Please run \"%s\" first.", color.Red("PandaBuilder setup"))
		c.result = false
		return
	}

	if err := c.solutionData.LoadFrom(c.pandaFile); err != nil {
		log.Fatalf("\n** error: fails in reading Pandfile: %s", color.Yellow(c.pandaFile))
		c.result = false
		return
	}

	if 

	c.solutionData.LoadFromLock(c.pandaLockFile)

	return
}

func (c *CommandBootstrap) Done() int {
	return CmdResultSuccess
}
