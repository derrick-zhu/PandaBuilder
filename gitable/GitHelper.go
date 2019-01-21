package gitable

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

type GitProtocol interface {
	RepoType() int
	URL() string
	REF() string
}

func GitFetch(aGit GitProtocol, remote string) bool {
	if aGit == nil {
		log.Println("** error: invalid git for fetch.")
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
		log.Fatalf("** error: fails in run command %s\n", err)
		return false
	}
	return true
}

func GitRetrieveCommitHash(aGit GitProtocol, ref string) string {
	if aGit == nil {
		log.Println("** error: invalid git for retrieve commit hash")
		return ""
	}

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
		log.Fatalf("** error: fails in run command %s\n", err)
		return ""
	}

	return strings.Trim(string(buf.Bytes()), " \n\t")
}

func GitClone(aGit GitProtocol, local string) bool {
	if aGit == nil {
		log.Println("** error: invalid git for cloning")
		return false
	}

	if len(local) == 0 {
		log.Println("** error: invalid git for cloning, local path is needed.")
		return false
	}

	if len(aGit.URL()) == 0 || len(aGit.REF()) == 0 {
		log.Println("** error: invalid git for cloning, information is not complete.")
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
