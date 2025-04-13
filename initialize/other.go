package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"os"
	"server/global"
	"server/utils"
)

// OtherInit 鎵ц鍏朵粬閰嶇疆鍒濆鍖?
func OtherInit() {
	// 瑙ｆ瀽鍒锋柊浠ょ墝杩囨湡鏃堕棿
	refreshTokenExpiry, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse refresh token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	// 瑙ｆ瀽璁块棶浠ょ墝杩囨湡鏃堕棿
	_, err = utils.ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse access token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	// 閰嶇疆鏈湴缂撳瓨杩囨湡鏃堕棿锛堜娇鐢ㄥ埛鏂颁护鐗岃繃鏈熸椂闂达紝鏂逛究鍦ㄨ繙绋嬬櫥褰曟垨璐︽埛鍐荤粨鏃跺 JWT 杩涜榛戝悕鍗曞鐞嗭級
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(refreshTokenExpiry),
	)
}