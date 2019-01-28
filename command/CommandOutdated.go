package command

import (
	"PandaBuilder/logger"
	"PandaBuilder/models"
	"path/filepath"
)

type CommandOutdated struct {
	solutionData  *models.PandaSolutionModel
	pandaFile     string
	pandaLockFile string
}

func NewCommandOutdated(argv []string) *CommandOutdated {
	result := &CommandOutdated{}
	result.Start(argv)
	return result
}

func (c *CommandOutdated) Start(argv []string) bool {
	if len(argv) < 3 {
		logger.Fatal("\n** Error: invalid parameters %v", argv)
		return false
	}

	c.pandaFile = argv[1]
	c.pandaLockFile = argv[2]
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

func (c *CommandOutdated) Run() {
	var err error
	if err = c.solutionData.LoadFrom(c.pandaFile); err != nil {
		logger.Fatal("\n** Error: fails in reading Pandfile: %v.", err)
		return
	}

	if lockResult := c.solutionData.LoadFromLock(c.pandaLockFile); lockResult != models.Success {
		logger.Fatal("\n** Error: fails in reading Pandfile.lock: %s.", c.pandaLockFile)
		return
	}

	if c.solutionData.LoadRemoteCommit() == false {
		logger.Fatal("\n** Error: fails in reading remote repo commit hash.")
		return
	}

	var modified bool
	if modified, err = c.solutionData.SyncLockData(); err != nil {
		logger.Fatal("\n** Error: fails in sync repo lock: %v.", err)
		return
	}

	if modified {
		if err = c.solutionData.SyncLockFile(c.pandaLockFile); err != nil {
			logger.Fatal("\n** Error: fails in sync repo lock file: %v.", err)
			return
		}
		logger.Println("Info: Panda lock file had been updated.")
	} else {
		logger.Println("Info: No commit updated. Skip write Panda lock file.")
	}
}

func (c *CommandOutdated) Done() int {
	return CmdResultSuccess
}
