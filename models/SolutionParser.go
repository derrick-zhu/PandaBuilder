package models

import (
	"PandaBuilder/gitable"
	"PandaBuilder/logger"
	"PandaBuilder/yamlParser"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bclicn/color"
)

const (
	NotExisted = iota
	Success
	Existed
	OperatorFails
)

type PandaSolutionModel struct {
	Modules   *ModulesModel
	Libraries *LibrariesModel
	Locks     []*PandaSolutionLockModel
	Remote    []*PandaSolutionLockModel
}

func (c *PandaSolutionModel) ModuleWithName(name string) *ModuleModel {
	return c.Modules.ModuleWithName(name)
}

func (c *PandaSolutionModel) ModuleWithUrl(url string) *ModuleModel {
	if len(url) <= 0 {
		return nil
	}

	aLib := c.LibraryWithUrl(url)
	aMod := c.ModuleWithName(aLib.Name)
	if aLib == nil || aMod == nil {
		return nil
	}
	return aMod
}

func (c *PandaSolutionModel) LibraryWithIndex(index int) *LibraryModel {
	return c.Libraries.LibraryWithIndex(index)
}

func (c *PandaSolutionModel) LibraryWithName(name string) *LibraryModel {
	return c.Libraries.LibraryWithName(name)
}

func (c *PandaSolutionModel) LibraryWithUrl(url string) *LibraryModel {
	return c.Libraries.LibraryWithUrl(url)
}

func (c *PandaSolutionModel) NumOfLibraries() int {
	return c.Libraries.Length()
}

func (c *PandaSolutionModel) LockedLibraryWithUrl(url string) *PandaSolutionLockModel {
	if len(url) == 0 {
		return nil
	}
	for _, eachLock := range c.Locks {
		if eachLock.url == url {
			return eachLock
		}
	}
	return nil
}

func (c *PandaSolutionModel) AppendLockLibrary(lockedModule *PandaSolutionLockModel) {
	if lockedModule != nil {
		c.Locks = append(c.Locks, lockedModule)
	}
}

func (c *PandaSolutionModel) RemoteLibraryWithUrl(url string) *PandaSolutionLockModel {
	if len(url) == 0 {
		return nil
	}

	for _, eachRemote := range c.Remote {
		if eachRemote.url == url {
			return eachRemote
		}
	}
	return nil
}

func (self *PandaSolutionModel) SetupPandafile(path string) bool {
	if len(path) == 0 {
		logger.Fatal("\n** Error: invalid directory for the flutter project workspace.")
		return false
	}

	pandaFile := filepath.Join(path, "Pandafile")

	if fi, _ := os.Stat(pandaFile); fi != nil {
		logger.Println("warning: Pandafile is existed.\n** Skip setup flutter project environments.")
		return false
	}

	temp :=
		`modules:
#  - appA:
#    path: appA/
#    dependency:
#	  - libA
#  - libA:
#	  path: libA/
#	  dependency:
#		- libB
#  - libB:
#      path: appA/
#      dependency:

libraries:
#  - libA:
#      git:
#        url: xx/xx/xx.git
#		ref: develop
#  - libB:
#      git:
#        url: yy/yy/yy.git
#        ref: develop
	`

	if err := ioutil.WriteFile(pandaFile, bytes.NewBufferString(temp).Bytes(), os.ModePerm); err != nil {
		logger.Fatal("\n** Error: fails in write panda file: %s", pandaFile)
		return false
	} else {
		logger.Println("Flutter project environment had been setted up. \n** JOB DONE!. Have fun.")
		return true
	}
}

func (c *PandaSolutionModel) LoadFrom(path string) error {
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return errors.New("error: \"" + path + "\" is not existed ")
	}

	ce := yamlParser.ConfigEngine{}
	if err := ce.LoadFromYaml(path); err != nil {
		logger.Error("\n Error: \"%v\" in reading \"%s\".", err, path)
		return err
	}

	var idiot interface{}
	idiot = ce.Get("modules")
	c.Modules = &ModulesModel{}
	c.Modules.loadModules(idiot.([]interface{}))

	idiot = ce.Get("libraries")
	c.Libraries = &LibrariesModel{}
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
		logger.Fatal("\n** Error: error in reading %s: %v", lockFile, err)
		return OperatorFails
	}

	lockTextLines := strings.Split(string(buffer), "\n")
	c.Locks = []*PandaSolutionLockModel{}
	for _, eachLockLine := range lockTextLines {
		pLock := &PandaSolutionLockModel{}
		if false == pLock.LoadFromLock(eachLockLine) {
			log.Fatalf("\n** fatal: fails in reading lock item: %s", eachLockLine)
			break
		}
		c.Locks = append(c.Locks, pLock)
	}
	///
	return Success
}

func (c *PandaSolutionModel) LoadRemoteCommit() bool {
	if c.Remote == nil {
		c.Remote = []*PandaSolutionLockModel{}
	}

	for _, eachRepo := range c.Libraries.Libraries {
		logger.Printf("Fetching %s", color.BWhite(eachRepo.Name))
		curRepoRefCommitHash := gitable.GitRetrieveCommitHash(eachRepo.Git, eachRepo.Git.REF())
		if len(curRepoRefCommitHash) <= 0 {
			logger.Log("\n Warning: could not fetch library ref: %s's commit hash", eachRepo.Git.REF())
			continue
		}
		pRemote := NewPandaSolutionLockModel(
			eachRepo.Git.RepoType(),
			eachRepo.Git.URL(),
			eachRepo.Git.REF(),
			curRepoRefCommitHash,
		)
		c.Remote = append(c.Remote, pRemote)
		logger.PrintlnRaw(" -> GIT: %s, Ref: %s, Commit: %s", pRemote.URL(), pRemote.REF(), pRemote.CommitHash)
	}
	return true
}

func (c *PandaSolutionModel) SyncLockData() (bool, error) {
	var err error
	var modified bool

	for idx := 0; idx < c.Libraries.Length(); idx++ {
		libData := c.LibraryWithIndex(idx)
		lockedData := c.LockedLibraryWithUrl(libData.Git.url)
		remoteData := c.RemoteLibraryWithUrl(libData.Git.url)

		if remoteData == nil {
			logger.Println("warning: could not fetch \"%s\" remote commit.", libData.Name)
			err = fmt.Errorf("\n** Error: could not find remote library: %s [%s].", libData.Name, libData.Git.URL())
			break
		}

		if lockedData == nil {
			logger.Println("Update %s: NEW to commit: %s", libData.Name, remoteData.CommitHash)
			lockedData = NewPandaSolutionLockModel(libData.Git.repoType, libData.Git.url, libData.Git.ref, remoteData.CommitHash)
			c.AppendLockLibrary(lockedData)
			modified = true

		} else {
			if lockedData.CommitHash == remoteData.CommitHash {
				// SKIP, nonthing changed for this lib
				logger.Println("Using: %s: (%s)", color.BWhite(libData.Name), color.Green(lockedData.CommitHash))

			} else {
				logger.Println("Checkout %s from \"%s\" to \"%s\"", color.BWhite(libData.Name), color.Red(lockedData.CommitHash), color.Red(remoteData.CommitHash))
				lockedData.repoType = libData.Git.repoType
				lockedData.ref = libData.Git.ref
				lockedData.CommitHash = remoteData.CommitHash
				modified = true
			}
		}

		if err != nil {
			break
		}
	}
	return modified, err
}

func (c *PandaSolutionModel) SyncLockFile(lockFile string) error {
	if len(lockFile) == 0 {
		logger.Error("\n Error: invalid lock file to write lock data")
		return fmt.Errorf("\n** error: invalid lock file to write lock data")
	}

	var err error
	allLockDescriptions := []string{}
	for _, eachLockData := range c.Locks {
		lockDescription := eachLockData.ToLockDescription()
		allLockDescriptions = append(allLockDescriptions, lockDescription)
	}

	result := strings.Join(allLockDescriptions, "\n")
	if err = ioutil.WriteFile(lockFile, bytes.NewBufferString(result).Bytes(), os.ModePerm); err != nil {
		return fmt.Errorf("\n** error: fails to write Pandafile.lock. %v", err)
	}

	return err
}

// PRIVATE HELPERS
