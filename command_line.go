package main

import (
	"flag"
)

const (
	setup     = 1 << iota // init
	update                // check and update local into latest commit
	outdated              // check and update Pandafile.lock
	bootstrap             // according Pandafile.lock, fetch
)

type CommandLine struct {
	Type   int
	Params []string
}

func (c *CommandLine) Parse() bool {
	cmdSetup := flag.Bool("setup", false, "init and setup flutter workspace")
	cmdUpdate := flag.Bool("update", false, "fetch and update Pandafile.lock and its flutter workspace")
	cmdOutdated := flag.Bool("outdated", false, "fetch and update Pandafile.lock")
	cmdBootstrap := flag.Bool("bootstrap", false, "update flutter workspace by Pandafile.lock")

	flag.Parse()

	//log.Printf("\nsetup:%t \nupdate:%t \noutdated:%t \nbootstrap:%t ", *cmdSetup, *cmdUpdate, *cmdOutdated, *cmdBootstrap)

	if *cmdBootstrap {
		c.Type = bootstrap
	} else if *cmdOutdated {
		c.Type = outdated
	} else if *cmdUpdate {
		c.Type = update
	} else if *cmdSetup {
		c.Type = setup
	}
	return true
}

func (c *CommandLine) Append(newParam string) {
	if c.Params == nil {
		c.Params = []string{}
	}
	c.Params = append(c.Params, newParam)
}
