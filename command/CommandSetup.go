package command

type CommandSetup struct {
}

func NewCommandSetup(argv []string) *CommandSetup {
	result := &CommandSetup{}
	result.Start(argv)
	return result
}

func (c *CommandSetup) Start(param []string) bool {
	return true
}

func (c *CommandSetup) Run() {
	return
}

func (c *CommandSetup) Done() int {
	return CmdResultSuccess
}
