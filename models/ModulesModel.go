package models

import (
	"PandaBuilder/logger"
	"bytes"
	"encoding/gob"
	"log"
)

// ----------------------------------------------------------------

type ModuleModel struct {
	name       string
	Path       string
	Dependency []string
}

func NewModuleModel(name string, path string, dependency []string) *ModuleModel {
	var err error
	result := &ModuleModel{}
	result.name = name
	result.Path = path

	if dependency == nil {
		dependency = []string{}
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err = enc.Encode(dependency); err != nil {
		logger.Error("\n** Error: fails in make module model[encode] %v.", err)
	}

	dec := gob.NewDecoder(&buf)
	if err = dec.Decode(&result.Dependency); err != nil {
		logger.Error("\n** Error: fails in make module model[decode] %v.", err)
	}

	return result
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
					logger.Log("\n Warning: mapping a undefined key: %s with value: %v", skey, sval)
				}
			}
		default:
			logger.Error("\n** Fails: in mapping data: %v into ModuleModel", value)
		}
		break // sorry, just only once, and who don't know about key's value, so using a iterator
	}
	return true
}

// ----------------------------------------------------------------

type ModulesModel struct {
	Modules []ModuleModel
}

func (c *ModulesModel) ModuleWithIndex(index int) *ModuleModel {
	if index < 0 || index >= len(c.Modules) {
		return nil
	}
	return &c.Modules[index]
}

func (c *ModulesModel) ModuleWithName(name string) *ModuleModel {
	for _, mod := range c.Modules {
		if mod.name == name {
			return &mod
		}
	}
	return nil
}

func (c *ModulesModel) NumOfModule() int {
	if c.Modules == nil {
		return 0
	}
	return len(c.Modules)
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
