package main

// import (
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"strings"
// )

// const (
// 	// Git define package type: Git
// 	Git = 1 << iota
// 	Unknown
// )

// // PackageModel define single Package info
// type PackageModel struct {
// 	Type   int
// 	URL    string
// 	Branch string
// 	Tag    string
// 	Local  string
// }

// func (c *PackageModel) parserType(pkgType string) {
// 	if strings.Trim(pkgType, " ") == "git" {
// 		c.Type = Git
// 	} else {
// 		c.Type = Unknown
// 	}
// }

// func (c *PackageModel) parserURL(url string) {
// 	url = strings.Trim(url, " ")
// 	if len(url) <= 0 {
// 		return
// 	}

// 	if strings.HasPrefix(url, "\"") && strings.HasSuffix(url, "\"") {
// 		url = strings.TrimPrefix(url, "\"")
// 		url = strings.TrimPrefix(url, "\"")
// 	}
// 	c.URL = url
// }

// func (c *PackageModel) parserOther(option string) {
// 	option = strings.Trim(option, " ")
// 	if len(option) <= 0 {
// 		return
// 	}
// 	var vals []string
// 	if vals = strings.Split(option, "=>"); len(vals) != 2 {
// 		return
// 	}

// 	dataType := strings.Trim(vals[0], " ")
// 	data := strings.Trim(vals[1], " ")

// 	if dataType == ":branch" {
// 		c.Branch = data
// 	} else if dataType == ":tag" {
// 		c.Tag = data
// 	} else if dataType == ":local" {
// 		c.Local = data
// 	} else {
// 		// EMPTY
// 	}
// }

// func (c *PackageModel) cloneGit() bool {
// 	if len(c.URL) <= 0 {
// 		return false
// 	}

// 	cmd := []string{}
// 	cmd = append(cmd, "clone")
// 	cmd = append(cmd, c.URL)
// 	cmd = append(cmd, "--single-branch")
// 	cmd = append(cmd, "--branch")

// 	if c.Type == Git {
// 		if len(c.Tag) > 0 {
// 			cmd = append(cmd, c.Tag)
// 		} else if len(c.Branch) > 0 {
// 			cmd = append(cmd, c.Branch)
// 		} else {
// 			// EMPTY
// 		}

// 		if len(c.Local) > 0 {
// 			cmd = append(cmd, c.Local)
// 		}
// 	}

// 	shellcmd := exec.Command("git", cmd...)
// 	shellcmd.Stdout = os.Stdout
// 	shellcmd.Stdin = os.Stdin
// 	shellcmd.Stderr = os.Stderr

// 	// fmt.Printf("%v", shellcmd)
// 	if err := shellcmd.Run(); err != nil {
// 		log.Fatalf("fails in cloning git resposity: %s. (%v)", c.URL, err)
// 		return false
// 	}
// 	return true
// }

// // SolutionPackageModel define solution package dependency info
// type SolutionPackageModel struct {
// 	Package []PackageModel
// }

// func (c *SolutionPackageModel) loadSolutionPackage(file string) {
// 	c.Package = []PackageModel{}

// 	var slnPkgBuff []byte
// 	var err error

// 	if slnPkgBuff, err = ioutil.ReadFile(file); err != nil {
// 		log.Fatalf("fails in reading '%s': %v", file, err)
// 		return
// 	}

// 	slnPkg := string(slnPkgBuff)
// 	pkgs := strings.Split(slnPkg, "\n")

// 	for _, eachPkg := range pkgs {
// 		if len(strings.Trim(eachPkg, " ")) <= 0 {
// 			continue
// 		}

// 		aPkg := PackageModel{}
// 		subPkg := strings.Split(eachPkg, ",")

// 		aPkg.parserType(subPkg[0])
// 		aPkg.parserURL(subPkg[1])
// 		for idx := 2; idx < len(subPkg); idx++ {
// 			aPkg.parserOther(subPkg[idx])
// 		}
// 		c.Package = append(c.Package, aPkg)
// 	}
// }
