package gitable

import (
	"PandaBuilder/logger"
	"PandaBuilder/shell"
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// GitProtocol define the rule of the GIT common protocol.
type GitProtocol interface {
	RepoType() int
	URL() string
	REF() string
}

// GitFetch fetch the remote changment
func GitFetch(aGit GitProtocol, remote string) bool {
	if aGit == nil {
		logger.Error("\n Error: invalid git for fetch.")
		return false
	}

	if len(remote) <= 0 {
		remote = "--all"
	}

	cmd := fmt.Sprintf("git fetch %s", remote)
	shellcmd := shell.ShellCommandWith(cmd)

	if err := shellcmd.Run(); err != nil {
		log.Fatalf("** error: fails in run command %s", err)
		return false
	}

	return true
}

// GitRetrieveCommitHash retrieve ref's commit hash of the GIT repo.
func GitRetrieveCommitHash(aGit GitProtocol, ref string) string {
	if aGit == nil {
		logger.Error("\n Error: invalid git for retrieve commit hash")
		return ""
	}

	if len(ref) == 0 {
		ref = "HEAD"
	}
	var buf bytes.Buffer

	cmdStr := fmt.Sprintf("git ls-remote %s %s", aGit.URL(), ref)
	shellcmd := shell.ShellCommandWith(cmdStr)
	shellcmd.Stdout = &buf

	if err := shellcmd.Run(); err != nil {
		log.Fatalf("** error: fails in run command %s", err)
		return ""
	}

	rawResult := strings.Trim(string(buf.Bytes()), "\t\n ")
	re := regexp.MustCompile(`(?m)(^[a-z0-9]+)[\t ]+(refs\/heads\/(\w+)|HEAD)`)
	filtedResult := re.FindAllString(rawResult, -1)

	if len(filtedResult) == 0 {
		log.Fatalf("** error: invalid commit hash fetched from remote repo: %s", aGit.URL())
	}

	filtedResult = strings.Fields(filtedResult[0])

	result := ""
	for index, eachValue := range filtedResult {
		if index%2 == 0 {
			continue
		}

		founded := false
		valueComponents := strings.Split(eachValue, "/")
		for _, eachComp := range valueComponents {
			founded = (eachComp == ref)
			if founded {
				break
			}
		}

		if founded {
			result = filtedResult[index-1]
		}
	}

	return result
}

func GetRetrieveCurrentGitCommitHash(ref string) string {

	if len(ref) == 0 {
		ref = "HEAD"
	}
	cmdStr := fmt.Sprintf("git rev-parse --verify %s", ref)
	shellcmd := shell.ShellCommandWith(cmdStr)

	var buf bytes.Buffer
	shellcmd.Stdout = &buf

	if err := shellcmd.Run(); err != nil {
		log.Fatalf("** error: fails in run command %s", err)
	}

	return strings.Trim(string(buf.Bytes()), " \n\t")
}

func GitClone(aGit GitProtocol, local string) bool {
	if aGit == nil {
		logger.Error("\n Error: invalid git for cloning")
		return false
	}

	if len(local) == 0 {
		logger.Error("\n Error: invalid git for cloning, local path is needed.")
		return false
	}

	if len(aGit.URL()) == 0 || len(aGit.REF()) == 0 {
		logger.Error("\n Error: invalid git for cloning, information is not complete.")
		return false
	}

	cmdStr := fmt.Sprintf("git clone %s --single-branch --branch %s %s", aGit.URL(), aGit.REF(), local)
	shellcmd := shell.ShellCommandWith(cmdStr)

	logger.Println("Clone %s...", aGit.URL())
	if err := shellcmd.Run(); err != nil {
		log.Fatalf("fails in cloning git resposity: %s. (%v)", aGit.URL(), err)
		return false
	}

	return true
}
