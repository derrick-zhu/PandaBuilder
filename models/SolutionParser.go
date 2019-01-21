package models

import (
	"PandaBuilder/yamlParser"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ModuleModel struct {
	name       string
	Path       string
	Dependency []string
}

func (c *ModuleModel) loadModule(mapValue map[interface{}]interface{}) bool {
	if mapValue == nil {
		log.Printf("warning: null data for ModuleModel")
		return true
	}

	for key, value := range mapValue {
		switch value.(type) {
		case map[interface{}]interface{}:
			c.name = key.(string)
			for skey, sval := range value.(map[interface{}]interface{}) {
				if skey == "path" {
					c.Path = sval.(string)
				} else if skey == "dependency" {
					c.Dependency = []string{}
					for _, dep := range sval.([]interface{}) {
						c.Dependency = append(c.Dependency, dep.(string))
					}
				} else {
					log.Printf("** warning: mapping a undefined key: %s with value: %v", skey, sval)
				}
			}
		default:
			log.Printf("** fails: in mapping data: %v into ModuleModel", value)
		}
		break // sorry, just only once, and who don't know about key's value, so using a iterator
	}
	return true
}

type ModulesModel struct {
	Modules []ModuleModel
}

func (c *ModulesModel) ModuleWith(name string) *ModuleModel {
	for _, mod := range c.Modules {
		if mod.name == name {
			return &mod
		}
	}
	return nil
}

func (c *ModulesModel) loadModules(values []interface{}) bool {
	if values == nil {
		return false
	}

	c.Modules = []ModuleModel{}
	for _, value := range values {
		moduleValue := ModuleModel{}
		moduleValue.loadModule(value.(map[interface{}]interface{}))
		c.Modules = append(c.Modules, moduleValue)
	}
	return true
}

const (
	Git = 1 << iota
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
		log.Printf("** waring: unknow repo type \"%s\"", repo)
	}
}

func (c *PackageGitModel) loadGitData(data map[interface{}]interface{}) bool {
	if data == nil {
		log.Printf("** warning: null data for PackageGitModel")
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
					log.Printf("** warnning: mapping a undefined key: %s with value: %v", propKey, propValue)
				}
			}
		default:
			log.Printf("** fails: in mapping data: %v into PackageGitModel", value)
		}
		break
	}

	return true
}

const (
	NotExisted = iota
	Success
	Existed
	OperatorFails
)

type PandaSolutionLockModel struct {
	PackageGitModel
	lockRef string
}

func (c *PandaSolutionLockModel) loadFrom(pandaLine string) bool {
	if len(pandaLine) <= 0 {
		log.Printf("** warning: panda lock text is empty.")
		return false
	}

	var allSegments []string = strings.Split(pandaLine, " ")
	if len(allSegments) != 3 {
		log.Printf("** warning: invalid lock item %s", pandaLine)
		return false
	}

	for idx, eachSegment := range allSegments {
		val := strings.Trim(eachSegment, " \"")
		if idx == 0 {
			c.SetUpWith(val)
		} else if idx == 1 {
			c.url = val
		} else if idx == 2 {
			c.lockRef = val
		} else {
			// EMPTY
		}
	}

	return true
}

type LibraryModel struct {
	Name string
	Git  PackageGitModel
}

func (c *LibraryModel) loadLibrary(data map[interface{}]interface{}) bool {
	if data == nil {
		log.Printf("** warning: null data for LibraryModel")
		return true
	}

	for key, value := range data {
		switch value.(type) {
		case map[interface{}]interface{}:
			c.Name = key.(string)
			c.Git = PackageGitModel{}
			c.Git.loadGitData(value.(map[interface{}]interface{}))

		default:
			log.Printf("** warning: invalid type for %s of data: %v", key, value)
		}
		break
	}
	return true
}

type LibrariesModel struct {
	Libraries []LibraryModel
}

func (c *LibrariesModel) LibraryWithIndex(index int) *LibraryModel {
	if index < 0 || index >= len(c.Libraries) {
		return nil
	}
	return &c.Libraries[index]
}

func (c *LibrariesModel) Length() int {
	return len(c.Libraries)
}

func (c *LibrariesModel) loadLibraries(libs []interface{}) bool {
	if libs == nil {
		return false
	}

	c.Libraries = []LibraryModel{}
	for _, value := range libs {
		aLib := LibraryModel{}
		aLib.loadLibrary(value.(map[interface{}]interface{}))
		c.Libraries = append(c.Libraries, aLib)
	}
	return true
}

type PandaSolutionModel struct {
	Modules   ModulesModel
	Libraries LibrariesModel
	Locks     PandaSolutionLockModel
}

func (c *PandaSolutionModel) ModuleWith(name string) *ModuleModel {
	return c.Modules.ModuleWith(name)
}

func (c *PandaSolutionModel) LibraryWithIndex(index int) *LibraryModel {
	return c.Libraries.LibraryWithIndex(index)
}

func (c *PandaSolutionModel) NumOfLibraries() int {
	return c.Libraries.Length()
}

func (c *PandaSolutionModel) LoadFrom(path string) error {
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return errors.New("error: \"" + path + "\" is not existed ")
	}

	ce := yamlParser.ConfigEngine{}
	if ce.LoadFromYaml(path) != nil {
		return errors.New("error: fails in loading Pandafile: " + path)
	}

	var idiot interface{}
	idiot = ce.Get("modules")
	c.Modules = ModulesModel{}
	c.Modules.loadModules(idiot.([]interface{}))

	idiot = ce.Get("libraries")
	c.Libraries = LibrariesModel{}
	c.Libraries.loadLibraries(idiot.([]interface{}))

	return nil
}

func (c *PandaSolutionModel) LoadFromLock(lockFile string) int {
	var err error
	if _, err = os.Stat(lockFile); os.IsNotExist(err) {
		return NotExisted
	}
	err = nil

	var buffer []byte
	if buffer, err = ioutil.ReadFile(lockFile); err != nil {
		log.Fatalf("** error: error in reading %s: %v", lockFile, err)
		return OperatorFails
	}

	lockTextLines := strings.Split(string(buffer), "\n")
	c.Locks = PandaSolutionLockModel{}
	for _, eachLockLine := range lockTextLines {
		if false == c.Locks.loadFrom(eachLockLine) {
			log.Fatalf("** fatal: fails in reading lock item: %s", eachLockLine)
			break
		}
	}
	///
	return Success
}
