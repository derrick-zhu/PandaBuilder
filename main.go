package main

import (
	"PandaBuilder/gitable"
	"PandaBuilder/models"
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
	// yamlFile := execPath + "/pubspec.yaml"
	// newYamlFile := execPath + "/pubspec.modified.yaml"
	pandaFile := execPath + "/Pandafile"
	pandaLockFile := execPath + "/Pandafile.lock"

	// ce := ConfigEngine{}
	// ce.loadFromYaml(yamlFile)

	// res := ce.Get("dependencies")

	// mapRes := res.(map[interface{}]interface{})
	// mapRes["hello"] = "world"

	// fmt.Println(res)

	// ce.saveToYaml(newYamlFile)

	slnData := models.PandaSolutionModel{}
	slnData.LoadFrom(pandaFile)

	for idx := 0; idx < slnData.NumOfLibraries(); idx++ {
		aLib := slnData.LibraryWithIndex(idx)
		gitable.GitClone(aLib.Git, aLib.Name)
	}

}
