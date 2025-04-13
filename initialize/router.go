package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
)

// InitRouter 鍒濆鍖栬矾鐢?
func InitRouter() *gin.Engine {
	// 璁剧疆gin妯″紡
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	// TODO
	return Router
}