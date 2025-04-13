//go:build  windows
// +build windows

package core

import (
    "github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// initServer 鍑芥暟鍒濆鍖栦竴涓爣鍑嗙殑 HTTP 鏈嶅姟鍣紙閫傜敤浜?Windows 绯荤粺锛?
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,          // 璁剧疆鏈嶅姟鍣ㄧ洃鍚殑鍦板潃
		Handler:        router,           // 璁剧疆璇锋眰澶勭悊鍣紙璺敱锛?
		ReadTimeout:    10 * time.Minute, // 璁剧疆璇锋眰鐨勮鍙栬秴鏃舵椂闂翠负 10 鍒嗛挓
		WriteTimeout:   10 * time.Minute, // 璁剧疆鍝嶅簲鐨勫啓鍏ヨ秴鏃舵椂闂翠负 10 鍒嗛挓
		MaxHeaderBytes: 1 << 20,          // 璁剧疆鏈€澶ц姹傚ご鐨勫ぇ灏忥紙1MB锛?
	}
}