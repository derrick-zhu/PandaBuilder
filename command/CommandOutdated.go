package command

import (
	"PandaBuilder/models"
	"fmt"
	"log"
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
		log.Fatalf("** error: invalid parameters %v", argv)
		return false
	}

	c.pandaFile = argv[1]
	c.pandaLockFile = argv[2]
	c.solutionData = &models.PandaSolutionModel{}

	return true
}

func (c *CommandOutdated) Run() {
	var err error
	if err = c.solutionData.LoadFrom(c.pandaFile); err != nil {
		log.Fatalf("\n** error: fails in reading Pandfile: %s.", c.pandaFile)
		return
	}

	if lockResult := c.solutionData.LoadFromLock(c.pandaLockFile); lockResult != models.Success {
		log.Fatalf("\n** error: fails in reading Pandfile.lock: %s.", c.pandaLockFile)
		return
	}

	if c.solutionData.LoadRemoteCommit() == false {
		log.Fatalf("\n** error: fails in reading remote repo commit hash.")
		return
	}

	var modified bool
	if modified, err = c.solutionData.SyncLockData(); err != nil {
		log.Fatalf("\n** error: fails in sync repo lock: %v.", err)
		return
	}

	if modified {
		if err = c.solutionData.SyncLockFile(c.pandaLockFile); err != nil {
			log.Fatalf("\n** error: fails in sync repo lock file: %v.", err)
			return
		}
		fmt.Printf("\n** Info: Panda lock file had been updated.")
	} else {
		fmt.Printf("\n** Info: No commit updated. Skip write Panda lock file")
	}
}

func (c *CommandOutdated) Done() int {
	return CmdResultSuccess
}
