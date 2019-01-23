package command

type CommandUpdate struct {
}

func NewCommandUpdate(argv []string) *CommandUpdate {
	result := &CommandUpdate{}
	result.Start(argv)
	return result
}

func (c *CommandUpdate) Start(param []string) bool {
	return true
}

func (c *CommandUpdate) Run() {
	return
}

func (c *CommandUpdate) Done() int {
	return CmdResultSuccess
}
