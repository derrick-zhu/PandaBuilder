package models

import (
	"PandaBuilder/logger"
	"log"
	"strings"
)

const (
	// Git package repo type: git repo
	Git = iota
)

// PackageGitModel the struct about the information of the git repo
type PackageGitModel struct {
	repoType int
	url      string
	ref      string
}

// RepoType repo type in integer format
func (c PackageGitModel) RepoType() int {
	return c.repoType
}

// URL repo's url address
func (c PackageGitModel) URL() string {
	return c.url
}

// REF current repo state, it could be a branch name, a tag or a commit hash(really??)
func (c PackageGitModel) REF() string {
	return c.ref
}

// SetUpWith set with the repo type by string, it could be git, etc.
func (c *PackageGitModel) SetUpWith(repo string) {
	if strings.ToLower(repo) == "git" {
		c.repoType = Git
	} else {
		log.Printf("\n** waring: unknow repo type \"%s\"", repo)
	}
}

// RepoTypeDescription get the repo type in string format, it could returns git etc.
func (c *PackageGitModel) RepoTypeDescription() string {
	Repo := []string{"git"}
	if c.repoType < 0 || c.repoType >= len(Repo) {
		logger.Fatal("\n** Error: invalid repo type: %d", c.repoType)
		return ""
	}
	return Repo[c.repoType]
}

func (c *PackageGitModel) loadGitData(data map[interface{}]interface{}) bool {
	if data == nil {
		logger.Log("\n Warning: null data for PackageGitModel")
		return true
	}

	for key, value := range data {
		switch value.(type) {
		case map[interface{}]interface{}:
			c.SetUpWith(key.(string))

			for propKey, propValue := range value.(map[interface{}]interface{}) {
				if strings.ToLower(propKey.(string)) == "url" {
					c.url = propValue.(string)
				} else if strings.ToLower(propKey.(string)) == "ref" {
					c.ref = propValue.(string)
				} else {
					logger.Verbose("\n** Warnning: mapping a undefined key: %s with value: %v", propKey, propValue)
				}
			}
		default:
			logger.Error("\n** Fails: in mapping data: %v into PackageGitModel", value)
		}
		break
	}

	return true
}
