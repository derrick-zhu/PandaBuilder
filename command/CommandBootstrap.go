package command

type CommandBootstrap struct {
}

func NewCommandBootstrap(argv []string) *CommandBootstrap {
	result := &CommandBootstrap{}
	result.Start(argv)
	return result
}

func (c *CommandBootstrap) Start(param []string) bool {
	return true
}

func (c *CommandBootstrap) Run() {
	return
}

func (c *CommandBootstrap) Done() int {
	return CmdResultSuccess
}
