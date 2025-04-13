package config

// Jwt jwt 閰嶇疆
type Jwt struct {
	AccessTokenSecret      string `json:"access_token_secret" yaml:"access_token_secret"`             // 鐢ㄤ簬鐢熸垚鍜岄獙璇佽闂护鐗岀殑瀵嗛挜
	RefreshTokenSecret     string `json:"refresh_token_secret" yaml:"refresh_token_secret"`           // 鐢ㄤ簬鐢熸垚鍜岄獙璇佸埛鏂颁护鐗岀殑瀵嗛挜
	AccessTokenExpiryTime  string `json:"access_token_expiry_time" yaml:"access_token_expiry_time"`   // 璁块棶浠ょ墝鐨勮繃鏈熸椂闂达紝渚嬪 "15m" 琛ㄧず 15 鍒嗛挓
	RefreshTokenExpiryTime string `json:"refresh_token_expiry_time" yaml:"refresh_token_expiry_time"` // 鍒锋柊浠ょ墝鐨勮繃鏈熸椂闂达紝渚嬪 "30d" 琛ㄧず 30 澶?
	Issuer                 string `json:"issuer" yaml:"issuer"`                                       // JWT 鐨勭鍙戣€呬俊鎭紝閫氬父鏄簲鐢ㄦ垨鏈嶅姟鐨勫悕绉?
}