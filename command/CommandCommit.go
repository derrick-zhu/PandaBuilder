package command

import "PandaBuilder/models"

type CommandCommit struct {
	solutionData  *models.PandaSolutionModel
	solutionPath  string
	pandaFile     string
	pandaLockFile string
	result        bool
}

func NewCommandCommit(argv []string) *CommandCommit {
	result := &CommandCommit{}
	result.Start(argv)
	return result
}

func (c *CommandCommit) Start(param []string) bool {
	return true
}

func (c *CommandCommit) Run() {
	return
}

func (c *CommandCommit) Done() int {
	return 0
}
