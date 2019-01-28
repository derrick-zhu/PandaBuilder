package models

import "PandaBuilder/logger"

type LibraryModel struct {
	Name string
	Git  PackageGitModel
}

func (c *LibraryModel) loadLibrary(data map[interface{}]interface{}) bool {
	if data == nil {
		logger.Log("\n Warning: null data for LibraryModel")
		return true
	}

	for key, value := range data {
		switch value.(type) {
		case map[interface{}]interface{}:
			c.Name = key.(string)
			c.Git = PackageGitModel{}
			c.Git.loadGitData(value.(map[interface{}]interface{}))

		default:
			logger.Log("\n Warning: invalid type for %s of data: %v", key, value)
		}
		break
	}
	return true
}

type LibrariesModel struct {
	Libraries []*LibraryModel
}

func (c *LibrariesModel) LibraryWithIndex(index int) *LibraryModel {
	if index < 0 || index >= len(c.Libraries) {
		return nil
	}
	return c.Libraries[index]
}

func (c *LibrariesModel) LibraryWithName(name string) *LibraryModel {
	if len(name) <= 0 {
		return nil
	}
	for _, eachLib := range c.Libraries {
		if eachLib.Name == name {
			return eachLib
		}
	}
	return nil
}

func (c *LibrariesModel) LibraryWithUrl(url string) *LibraryModel {
	if len(url) <= 0 {
		return nil
	}

	for _, eachLib := range c.Libraries {
		if eachLib.Git.URL() == url {
			return eachLib
		}
	}
	return nil
}

func (c *LibrariesModel) Length() int {
	return len(c.Libraries)
}

func (c *LibrariesModel) loadLibraries(libs []interface{}) bool {
	if libs == nil {
		return false
	}

	c.Libraries = []*LibraryModel{}
	for _, value := range libs {
		aLib := &LibraryModel{}
		aLib.loadLibrary(value.(map[interface{}]interface{}))
		c.Libraries = append(c.Libraries, aLib)
	}
	return true
}
