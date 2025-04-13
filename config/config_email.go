package config

// Email 閭閰嶇疆锛屽彲浠ョ櫥褰?QQ 閭锛屾墦寮€璁剧疆锛屽紑鍚涓夋柟鏈嶅姟鏈嶅姟锛岃鎯呰瑙?https://mail.qq.com/
type Email struct {
	Host     string `json:"host" yaml:"host"`         // 閭欢鏈嶅姟鍣ㄥ湴鍧€锛屼緥濡?smtp.qq.com
	Port     int    `json:"port" yaml:"port"`         // 閭欢鏈嶅姟鍣ㄧ鍙ｏ紝甯歌鐨勫 587 (TLS) 鎴?465 (SSL)
	From     string `json:"from" yaml:"from"`         // 鍙戜欢浜洪偖绠卞湴鍧€
	Nickname string `json:"nickname" yaml:"nickname"` // 鍙戜欢浜烘樀绉帮紝鐢ㄤ簬鏄剧ず鍦ㄩ偖浠朵腑鐨勫彂浠朵汉淇℃伅
	Secret   string `json:"secret" yaml:"secret"`     // 鍙戜欢浜洪偖绠辩殑瀵嗙爜鎴栧簲鐢ㄤ笓鐢ㄥ瘑鐮侊紝鐢ㄤ簬韬唤楠岃瘉
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl"`     // 鏄惁浣跨敤 SSL 鍔犲瘑杩炴帴锛宼rue 琛ㄧず浣跨敤锛宖alse 琛ㄧず涓嶄娇鐢?
}