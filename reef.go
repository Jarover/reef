package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/Jarover/reef/readconfig"
	"github.com/Jarover/reef/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Читаем флаги и окружение
func readFlag(configFlag *readconfig.Flag) {
	flag.StringVar(&configFlag.ConfigFile, "f", readconfig.GetEnv("CONFIGFILE", utils.GetBaseFile()+".json"), "config file")
	//flag.StringVar(&configFlag.Host, "h", readconfig.GetEnv("HOST", ""), "host")
	flag.UintVar(&configFlag.Port, "p", uint(readconfig.GetEnvInt("PORT", 0)), "port")
	flag.Parse()

}

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
	logPath := dir + "/" + utils.GetBaseFile() + "_app.log"
	fmt.Println(logPath)
	l := &lumberjack.Logger{ //nolint:typecheck
		Filename:   logPath,
		MaxSize:    500, // megabytes
		MaxBackups: 10,
		MaxAge:     1,     //days
		Compress:   false, // disabled by default
	}
	log.SetOutput(l)
	log.Println("Start!")
	r := gin.Default()
	r.LoadHTMLGlob(dir + "/templates/*")
	r.GET("/", homePage)
	r.GET("/info", infoPage)
	r.GET("/api/levels", levels)

	r.Run(":" + strconv.FormatUint(uint64(Config.Port), 10)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
