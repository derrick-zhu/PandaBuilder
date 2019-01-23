package command

import "strings"

const (
	// CmdResultUnknown for unknown error about command execute result
	CmdResultUnknown = -1 << iota
	// CmdResultError for common error about command execute result
	CmdResultError = 0 << iota
	// CmdResultSuccess for success result about command execute result
	CmdResultSuccess
)

type CommandProtocol interface {
	Start(argv []string) bool
	Run()
	Done() int
}

func Factory(cmdType string, argv []string) CommandProtocol {
	tmp := strings.ToLower(cmdType)
	if tmp == "setup" {
		return NewCommandSetup(argv)
	} else if tmp == "outdated" {
		return NewCommandOutdated(argv)
	} else if tmp == "update" {
		return NewCommandUpdate(argv)
	} else if tmp == "bootstrap" {
		return NewCommandBootstrap(argv)
	} else {
		return nil
	}
}
