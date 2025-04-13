package config

// Redis 缂撳瓨鏁版嵁搴撻厤缃?
type Redis struct {
	Address  string `json:"address" yaml:"address"`   // Redis 鏈嶅姟鍣ㄧ殑鍦板潃锛岄€氬父涓?"localhost:6379" 鎴栧叾浠栦富鏈哄拰绔彛
	Password string `json:"password" yaml:"password"` // 杩炴帴 Redis 鏃剁殑瀵嗙爜锛屽鏋滄病鏈夎缃瘑鐮佸垯鐣欑┖
	DB       int    `json:"db" yaml:"db"`             // 鎸囧畾浣跨敤鐨勬暟鎹簱绱㈠紩锛屽崟瀹炰緥妯″紡涓嬪彲閫夋嫨鐨勬暟鎹簱锛岄粯璁や负 0
}