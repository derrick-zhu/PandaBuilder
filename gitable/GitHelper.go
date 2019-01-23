package gitable

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type GitProtocol interface {
	RepoType() int
	URL() string
	REF() string
}

func GitFetch(aGit GitProtocol, remote string) bool {
	if aGit == nil {
		log.Printf("\n** error: invalid git for fetch.")
		return false
	}
	cmd := []string{}
	cmd = append(cmd, "fetch")
	if len(remote) > 0 {
		cmd = append(cmd, remote)
	} else {
		cmd = append(cmd, "--all")
	}

	shellcmd := exec.Command("git", cmd...)
	var buf bytes.Buffer
	shellcmd.Stdout = &buf
	if err := shellcmd.Run(); err != nil {
		log.Fatalf("** error: fails in run command %s", err)
		return false
	}
	return true
}

func GitRetrieveCommitHash(aGit GitProtocol, ref string) string {
	if aGit == nil {
		log.Printf("\n** error: invalid git for retrieve commit hash")
		return ""
	}

	cmd := []string{}
	cmd = append(cmd, "ls-remote")
	cmd = append(cmd, aGit.URL())
	if len(ref) == 0 {
		ref = "HEAD"
	}
	cmd = append(cmd, ref)

	var buf bytes.Buffer

	shellcmd := exec.Command("git", cmd...)
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

	filtedResult = strings.Split(filtedResult[0], "\t")

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

	cmd := []string{}
	cmd = append(cmd, "rev-parse")
	if len(ref) == 0 {
		cmd = append(cmd, "HEAD")
	} else {
		cmd = append(cmd, ref)
	}

	shellcmd := exec.Command("git", cmd...)
	var buf bytes.Buffer
	shellcmd.Stdout = &buf
	if err := shellcmd.Run(); err != nil {
		log.Fatalf("** error: fails in run command %s", err)
	}

	return strings.Trim(string(buf.Bytes()), " \n\t")
}

func GitClone(aGit GitProtocol, local string) bool {
	if aGit == nil {
		log.Printf("\n** error: invalid git for cloning")
		return false
	}

	if len(local) == 0 {
		log.Printf("\n** error: invalid git for cloning, local path is needed.")
		return false
	}

	if len(aGit.URL()) == 0 || len(aGit.REF()) == 0 {
		log.Printf("\n** error: invalid git for cloning, information is not complete.")
		return false
	}

	cmd := []string{}
	cmd = append(cmd, "clone")
	cmd = append(cmd, aGit.URL())
	cmd = append(cmd, "--single-branch")
	cmd = append(cmd, "--branch")
	cmd = append(cmd, aGit.REF())
	cmd = append(cmd, local)

	shellcmd := exec.Command("git", cmd...)
	shellcmd.Stdout = os.Stdout
	shellcmd.Stdin = os.Stdin
	shellcmd.Stderr = os.Stderr

	log.Printf("Start... %v", shellcmd)
	if err := shellcmd.Run(); err != nil {
		log.Fatalf("fails in cloning git resposity: %s. (%v)", aGit.URL(), err)
		return false
	}

	return true
}
