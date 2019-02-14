package shell

import (
	"PandaBuilder/logger"
	"flag"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	Setup     = 1 << iota // init
	Update                // check and update local into latest commit
	Outdated              // check and update Pandafile.lock
	Bootstrap             // according Pandafile.lock, fetch
	Commit                // commit and push current modifications
)

type CommandLine struct {
	Type      int
	Workspace string
	Params    []string

	SkipUpdateLock       bool // 忽略更新Pandafile.lock
	SkipRepoClone        bool // 忽略clone各个库
	SkipUpdateSubPackage bool // 忽略更新各个库的pubspec.yaml
}

func (c *CommandLine) ShowHelp() {
	logger.Println("Usage: \n\tPandaBulder [workspace directory] [-setup|-update|-outdated|-bootstrap] <--skip-clone> <--no-update-lock> <--no-update-package>\n\nParameter:\n")
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
	cmdCommit := flag.Bool("commit", false, "commit the modification to remote repo.")

	// external flags
	c.SkipRepoClone = *flag.Bool("-skip-clone", false, "skip clone remote repo into local directory")
	c.SkipUpdateLock = *flag.Bool("-no-update-lock", false, "skip update current Pandafile.lock")
	c.SkipUpdateSubPackage = *flag.Bool("-no-update-package", false, "skip update each package's pubspec.yaml")

	flag.Parse()

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			continue
		}
		c.Workspace = arg
		break
	}

	if len(c.Workspace) <= 0 {
		c.Workspace = executePath()
	}

	if *cmdBootstrap {
		c.Type = Bootstrap
	} else if *cmdOutdated {
		c.Type = Outdated
	} else if *cmdUpdate {
		c.Type = Update
	} else if *cmdSetup {
		c.Type = Setup
	} else if *cmdCommit {
		c.Type = Commit
	}

	execPath := executePath()

	pandaFile := path.Join(c.Workspace, "Pandafile")
	pandaLockFile := path.Join(c.Workspace, "Pandafile.lock")

	c.AppendCommandParam(execPath)
	c.AppendCommandParam(pandaFile)
	c.AppendCommandParam(pandaLockFile)

	return true
}

func (c *CommandLine) AppendCommandParam(newParam string) {
	if c.Params == nil {
		c.Params = []string{}
	}
	c.Params = append(c.Params, newParam)
}

func executePath() string {
	var dir string
	var err error
	if dir, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		return ""
	}
	return dir
}
