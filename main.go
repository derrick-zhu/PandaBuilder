package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func executePath() string {
	var dir string
	var err error
	if dir, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		return ""
	}
	return dir
}

func main() {
	execPath := executePath()
	yamlFile := execPath + "/pubspec.yaml"
	newYamlFile := execPath + "/pubspec.modified.yaml"

	ce := ConfigEngine{}
	ce.loadFromYaml(yamlFile)

	res := ce.Get("dependencies")

	mapRes := res.(map[interface{}]interface{})
	mapRes["hello"] = "world"

	fmt.Println(res)

	ce.saveToYaml(newYamlFile)
}
