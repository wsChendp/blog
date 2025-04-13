package config

// Gaode 楂樺痉鏈嶅姟閰嶇疆锛岃鎯呰瑙?https://lbs.amap.com/
type Gaode struct {
	Enable bool   `json:"enable" yaml:"enable"` // 鏄惁寮€鍚珮寰锋湇鍔★紝true 琛ㄧず鍚敤锛宖alse 琛ㄧず绂佺敤
	Key    string `json:"key" yaml:"key"`       // 楂樺痉鏈嶅姟鐨勫簲鐢ㄥ瘑閽ワ紝鐢ㄤ簬韬唤楠岃瘉鍜屾湇鍔¤闂?
}