package core

import (
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	// "server/service"
)

type server interface {
	ListenAndServe() error
}

// RunServer 鐢ㄤ簬鍚姩鏈嶅姟鍣?
func RunServer() {
	addr := global.Config.System.Addr()
	Router := initialize.InitRouter()

	// 鍔犺浇鎵€鏈夌殑 JWT 榛戝悕鍗曪紝瀛樺叆鏈湴缂撳瓨
	// TODO service.LoadAll()

    // 鍒濆鍖栨湇鍔″櫒骞跺惎鍔?
	s := initServer(addr, Router)
	global.Log.Info("server run success on ", zap.String("address", addr))
	global.Log.Error(s.ListenAndServe().Error())
}