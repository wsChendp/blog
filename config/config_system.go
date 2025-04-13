package config

import (
	"fmt"
	"server/model/appTypes"
	"strings"
)

// System 绯荤粺閰嶇疆
type System struct {
	Host           string `json:"-" yaml:"host"`                          // 鏈嶅姟鍣ㄧ粦瀹氱殑涓绘満鍦板潃锛岄€氬父涓?0.0.0.0 琛ㄧず鐩戝惉鎵€鏈夊彲鐢ㄥ湴鍧€
	Port           int    `json:"-" yaml:"port"`                          // 鏈嶅姟鍣ㄧ洃鍚殑绔彛鍙凤紝閫氬父鐢ㄤ簬 HTTP 鏈嶅姟
	Env            string `json:"-" yaml:"env"`                           // Gin 鐨勭幆澧冪被鍨嬶紝渚嬪 "debug"銆?release" 鎴?"test"
	RouterPrefix   string `json:"-" yaml:"router_prefix"`                 // API 璺敱鍓嶇紑锛岀敤浜庢瀯寤?API 璺緞
	UseMultipoint  bool   `json:"use_multipoint" yaml:"use_multipoint"`   // 鏄惁鍚敤澶氱偣鐧诲綍鎷︽埅锛岄槻姝㈠悓涓€璐︽埛鍦ㄥ涓湴鏂瑰悓鏃剁櫥褰?
	SessionsSecret string `json:"sessions_secret" yaml:"sessions_secret"` // 鐢ㄤ簬鍔犲瘑浼氳瘽鐨勫瘑閽ワ紝纭繚浼氳瘽鏁版嵁鐨勫畨鍏ㄦ€?
	OssType        string `json:"oss_type" yaml:"oss_type"`               // 瀵瑰簲鐨勫璞″瓨鍌ㄦ湇鍔＄被鍨嬶紝濡?"local" 鎴?"qiniu"
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s System) Storage() appTypes.Storage {
	switch strings.ToLower(s.OssType) {
	case "local", "Local":
		return appTypes.Local
	case "qiniu", "Qiniu":
		return appTypes.Qiniu
	default:
		return appTypes.Local
	}
}