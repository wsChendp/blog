//go:build !windows
// +build !windows

package core

import (
    "github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

// initServer 鍑芥暟鍒濆鍖栦竴涓?Endless 鏈嶅姟鍣紙閫傜敤浜庨潪 Windows 绯荤粺锛?
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router) // 浣跨敤 endless 鍖呭垱寤轰竴涓柊鐨?HTTP 鏈嶅姟鍣ㄥ疄渚?
	s.ReadHeaderTimeout = 10 * time.Minute  // 璁剧疆璇锋眰澶寸殑璇诲彇瓒呮椂鏃堕棿涓?10 鍒嗛挓
	s.WriteTimeout = 10 * time.Minute       // 璁剧疆鍝嶅簲鍐欏叆鐨勮秴鏃舵椂闂翠负 10 鍒嗛挓
	s.MaxHeaderBytes = 1 << 20              // 璁剧疆鏈€澶ц姹傚ご鐨勫ぇ灏忥紙1MB锛?

	return s // 杩斿洖鍒涘缓鐨勬湇鍔″櫒瀹炰緥
}