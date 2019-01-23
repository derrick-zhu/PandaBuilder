package main

import (
	"PandaBuilder/command"
	"log"
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
	cl := CommandLine{}
	cl.Parse()

	execPath := executePath()
	pandaFile := execPath + "/Pandafile"
	pandaLockFile := execPath + "/Pandafile.lock"

	cl.Append(execPath)
	cl.Append(pandaFile)
	cl.Append(pandaLockFile)

	var cmd command.CommandProtocol

	switch cl.Type {
	case setup:
		cmd = command.Factory("setup", cl.Params)
		break

	case outdated:
		cmd = command.Factory("outdated", cl.Params)
		break

	case update:
		cmd = command.Factory("update", cl.Params)
		break

	case bootstrap:
		cmd = command.Factory("bootstrap", cl.Params)
		break

	default:
		log.Printf("\n** warning: invalid command: \"%v\"", cl.Params)
	}

	cmd.Run()
	cmd.Done()

	// ce := ConfigEngine{}
	// ce.loadFromYaml(yamlFile)

	// res := ce.Get("dependencies")

	// mapRes := res.(map[interface{}]interface{})
	// mapRes["hello"] = "world"

	// fmt.Println(res)

	// ce.saveToYaml(newYamlFile)

	// slnData := models.PandaSolutionModel{}
	// slnData.LoadFrom(pandaFile)
	// slnData.LoadFromLock(pandaLockFile)
	// slnData.LoadRemoteCommit()

	// fmt.Printf("%v", slnData)

	// for idx := 0; idx < slnData.NumOfLibraries(); idx++ {
	// 	aLib := slnData.LibraryWithIndex(idx)
	// 	gitable.GitClone(aLib.Git, aLib.Name)
	// }

}
