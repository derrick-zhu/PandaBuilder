package gitable

import (
	"log"
	"os"
	"os/exec"
)

type GitProtocol interface {
	URL() string
	REF() string
}

func GitFetch(aGit GitProtocol, remote string) bool {
	return true
}

func GitRetrieveCommitHash(aGit GitProtocol, ref string) string {
	if aGit == nil {
		log.Println("** error: invalid git for cloning")
		return ""
	}

	if len(ref) == 0 {
		ref = "HEAD"
	}

	return ""
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
