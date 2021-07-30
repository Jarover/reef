package main

import "fmt"

func main() {

	dir := utils.GetDir()
	err := readconfig.Version.ReadVersionFile(dir + "/version.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(readconfig.Version)
	var configFlag readconfig.Flag
	readFlag(&configFlag)

	fmt.Println(configFlag)
	fmt.Println(dir + "/" + configFlag.ConfigFile)
	Config, err := readconfig.ReadConfig(dir + "/" + configFlag.ConfigFile)
	if configFlag.Port != 0 {
		fmt.Println(Config)
		Config.Port = configFlag.Port
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(Config)
}
