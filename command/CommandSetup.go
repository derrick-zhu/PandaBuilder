package command

import (
	"PandaBuilder/models"
	"log"
)

type CommandSetup struct {
	solutionData *models.PandaSolutionModel
	solutionPath string
	result       bool
}

func NewCommandSetup(argv []string) *CommandSetup {
	result := &CommandSetup{}
	result.Start(argv)
	return result
}

func (c *CommandSetup) Start(param []string) bool {
	if len(param) < 1 {
		log.Fatalf("** error: invalid parameters %v", param)
		return false
	}

	c.solutionPath = param[0]
	c.solutionData = &models.PandaSolutionModel{}

	return true
}

func (c *CommandSetup) Run() {
	c.result = c.solutionData.SetupPandafile(c.solutionPath)
	return
}

func (c *CommandSetup) Done() int {
	if c.result {
		return CmdResultSuccess
	} else {
		return CmdResultError
	}
}
