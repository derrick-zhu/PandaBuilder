package main

import (
	"PandaBuilder/command"
	"PandaBuilder/logger"
	"PandaBuilder/shell"
)

func main() {
	cl := shell.CommandLine{}
	cl.Parse()

	var cmd command.CommandProtocol

	switch cl.Type {
	case shell.Setup:
		cmd = command.Factory("setup", cl.Params)
		break

	case shell.Outdated:
		cmd = command.Factory("outdated", cl.Params)
		break

	case shell.Update:
		cmd = command.Factory("update", cl.Params)
		break

	case shell.Bootstrap:
		cmd = command.Factory("bootstrap", cl.Params)
		break

	case shell.Commit:
		cmd = command.Factory("commit", cl.Params)
		break

	default:
		logger.Println("warning: invalid command: \"%v\"\n", cl.Type)
		cl.ShowHelp()

		return
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
