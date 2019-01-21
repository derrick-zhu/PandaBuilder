package main

import (
	"flag"
	"log"
)

const (
	setup     = 1 << iota // init
	update                // check and update local into latest commit
	outdated              // check and update Pandafile.lock
	bootstrap             // according Pandafile.lock, fetch
)

type CommandLine struct {
	Type  int
	param []string
}

func (c *CommandLine) Parse() bool {
	cmdSetup := flag.String("setup", "", "init and setup flutter workspace")
	cmdUpdate := flag.String("update", "", "fetch and update Pandafile.lock and its flutter workspace")
	cmdOutdated := flag.String("outdated", "", "fetch and update Pandafile.lock")
	cmdBootstrap := flag.String("bootstrap", "", "update flutter workspace by Pandafile.lock")

	flag.Parse()

	log.Printf("\nsetup:%s \nupdate:%s \noutdated:%s \nbootstrap:%s \n", *cmdSetup, *cmdUpdate, *cmdOutdated, *cmdBootstrap)

	if len(*cmdBootstrap) == 0 {
		c.Type = bootstrap
	} else if len(*cmdOutdated) == 0 {
		c.Type = outdated
	} else if len(*cmdUpdate) == 0 {
		c.Type = update
	} else if len(*cmdSetup) == 0 {
		c.Type = setup
	}
	return true
}
