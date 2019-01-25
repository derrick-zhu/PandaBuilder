package shell

import (
	"flag"
	"fmt"
)

const (
	Setup     = 1 << iota // init
	Update                // check and update local into latest commit
	Outdated              // check and update Pandafile.lock
	Bootstrap             // according Pandafile.lock, fetch
)

type CommandLine struct {
	Type   int
	Params []string

	SkipUpdateLock       bool // 忽略更新Pandafile.lock
	SkipRepoClone        bool // 忽略clone各个库
	SkipUpdateSubPackage bool // 忽略更新各个库的pubspec.yaml
}

func (c *CommandLine) ShowHelp() {
	fmt.Printf("\nUsage: \n\tPandaBulder [-setup|-update|-outdated|-bootstrap] <--skip-clone> <--no-update-lock> <--no-update-package>\n\nParameter:\n")
	flag.PrintDefaults()
	// 	fmt.Printf(`

	// Usage:
	// 	PandaBulder [-setup|-update|-outdated|-bootstrap] <--skip-clone> <--no-update-lock> <--no-update-package>

	// Parameters:
	// 	-setup:
	// 		init and setup flutter workspace
	// 	-update:
	// 		fetch and update Pandafile.lock and its flutter workspace
	// 	-outdated:
	// 		fetch and update Pandafile.lock
	// 	-bootstrap
	// 		update flutter workspace by Pandafile.lock

	// 	--skip-clone:
	// 		skip clone remote repo into local directory
	// 	--no-update-lock:
	// 		skip update current Pandafile.lock
	// 	--no-update-package:
	// 		skip update each package's pubspec.yaml
	// `)
}

func (c *CommandLine) Parse() bool {
	// base operation flags
	cmdSetup := flag.Bool("setup", false, "init and setup flutter workspace")
	cmdUpdate := flag.Bool("update", false, "fetch and update Pandafile.lock and its flutter workspace")
	cmdOutdated := flag.Bool("outdated", false, "fetch and update Pandafile.lock")
	cmdBootstrap := flag.Bool("bootstrap", false, "update flutter workspace by Pandafile.lock")

	// external flags
	c.SkipRepoClone = *flag.Bool("-skip-clone", false, "skip clone remote repo into local directory")
	c.SkipUpdateLock = *flag.Bool("-no-update-lock", false, "skip update current Pandafile.lock")
	c.SkipUpdateSubPackage = *flag.Bool("-no-update-package", false, "skip update each package's pubspec.yaml")

	flag.Parse()

	//log.Printf("\nsetup:%t \nupdate:%t \noutdated:%t \nbootstrap:%t ", *cmdSetup, *cmdUpdate, *cmdOutdated, *cmdBootstrap)

	if *cmdBootstrap {
		c.Type = Bootstrap
	} else if *cmdOutdated {
		c.Type = Outdated
	} else if *cmdUpdate {
		c.Type = Update
	} else if *cmdSetup {
		c.Type = Setup
	}
	return true
}

func (c *CommandLine) AppendCommandParam(newParam string) {
	if c.Params == nil {
		c.Params = []string{}
	}
	c.Params = append(c.Params, newParam)
}
