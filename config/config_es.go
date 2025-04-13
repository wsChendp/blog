package config

// ES ElasticSearch 閰嶇疆
type ES struct {
	URL            string `json:"url" yaml:"url"`                           // Elasticsearch 鏈嶅姟鐨?URL锛屼緥濡?http://localhost:9200
	Username       string `json:"username" yaml:"username"`                 // 鐢ㄤ簬杩炴帴 Elasticsearch 鐨勭敤鎴峰悕
	Password       string `json:"password" yaml:"password"`                 // 鐢ㄤ簬杩炴帴 Elasticsearch 鐨勫瘑鐮?
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"` // 鏄惁鍦ㄦ帶鍒跺彴鎵撳嵃 Elasticsearch 璇彞锛宼rue 琛ㄧず鎵撳嵃锛宖alse 琛ㄧず涓嶆墦鍗?
}