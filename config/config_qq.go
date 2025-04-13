package config

// QQ qq 鐧诲綍閰嶇疆锛岃鎯呰瑙?https://connect.qq.com/
type QQ struct {
	Enable      bool   `json:"enable" yaml:"enable"`             // 鏄惁鍚敤 qq 鐧诲綍锛宼rue 琛ㄧず鍚敤锛宖alse 琛ㄧず绂佺敤
	AppID       string `json:"app_id" yaml:"app_id"`             // 搴旂敤 ID
	AppKey      string `json:"app_key" yaml:"app_key"`           // 搴旂敤瀵嗛挜
	RedirectURI string `json:"redirect_uri" yaml:"redirect_uri"` // 缃戠珯鍥炶皟鍩?
}

func (qq QQ) QQLoginURL() string {
	return "https://graph.qq.com/oauth2.0/authorize?" +
		"response_type=code&" +
		"client_id=" + qq.AppID + "&" +
		"redirect_uri=" + qq.RedirectURI
}