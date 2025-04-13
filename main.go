package main

import (
	"server/core"
	"server/flag"
	"server/global"
	"server/initialize"
)
func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	
	initialize.OtherInit()

	global.DB = initialize.InitGorm()
	global.ESClient = initialize.ConnectES()
	global.Redis = initialize.ConnectRedis()

	defer global.Redis.Close()

	flag.InitFlag()

	initialize.InitCron()
	
	core.RunServer()
}
