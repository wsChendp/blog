package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)
func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	
	initialize.OtherInit()

	global.DB = initialize.InitGorm()
	global.ES = initialize.ConnectES()
	global.Redis = initialize.ConnectRedis()

	defer global.Redis.Close()
	initialize.InitCron()

	core.RunServer()
}
