package models

import (
	"PandaBuilder/logger"
	"fmt"
	"strings"
)

type PandaSolutionLockModel struct {
	PackageGitModel
	CommitHash string
}

func NewPandaSolutionLockModel(repoType int, url string, ref string, commitHash string) *PandaSolutionLockModel {
	result := new(PandaSolutionLockModel)
	result.repoType = repoType
	result.url = url
	result.ref = ref
	result.CommitHash = commitHash
	return result
}

func (c *PandaSolutionLockModel) LoadFromLock(pandaLine string) bool {
	if len(pandaLine) <= 0 {
		logger.Log("\n Warning: panda lock text is empty.")
		return false
	}

	allSegments := strings.Fields(pandaLine)
	if len(allSegments) != 3 {
		logger.Log("\n Warning: invalid lock item %s", pandaLine)
		return false
	}

	for idx, eachSegment := range allSegments {
		val := strings.Trim(eachSegment, " \"")
		if idx == 0 {
			c.SetUpWith(val)
		} else if idx == 1 {
			c.url = val
		} else if idx == 2 {
			c.CommitHash = val
		} else {
			// EMPTY
		}
	}

	return true
}

func (c *PandaSolutionLockModel) ToLockDescription() string {
	return fmt.Sprintf("\"%s\"\t\"%s\"\t\"%s\"", c.RepoTypeDescription(), c.URL(), c.CommitHash)
}
