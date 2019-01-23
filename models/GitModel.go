package models

import (
	"log"
	"strings"
)

const (
	Git = iota
)

type PackageGitModel struct {
	repoType int
	url      string
	ref      string
}

func (c PackageGitModel) RepoType() int {
	return c.repoType
}

func (c PackageGitModel) URL() string {
	return c.url
}

func (c PackageGitModel) REF() string {
	return c.ref
}

func (c *PackageGitModel) SetUpWith(repo string) {
	if strings.ToLower(repo) == "git" {
		c.repoType = Git
	} else {
		log.Printf("\n** waring: unknow repo type \"%s\"", repo)
	}
}

func (c *PackageGitModel) RepoTypeDescription() string {
	Repo := []string{"git"}
	if c.repoType < 0 || c.repoType >= len(Repo) {
		log.Fatalf("\n** error: invalid repo type: %d", c.repoType)
		return ""
	}
	return Repo[c.repoType]
}

func (c *PackageGitModel) loadGitData(data map[interface{}]interface{}) bool {
	if data == nil {
		log.Printf("\n** warning: null data for PackageGitModel")
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
					log.Printf("\n** warnning: mapping a undefined key: %s with value: %v", propKey, propValue)
				}
			}
		default:
			log.Printf("\n** fails: in mapping data: %v into PackageGitModel", value)
		}
		break
	}

	return true
}
