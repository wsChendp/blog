package config

import (
	"gorm.io/gorm/logger"
	"strconv"
	"strings"
)

// Mysql 鏁版嵁搴撻厤缃?
type Mysql struct {
	Host         string `json:"host" yaml:"host"`                     // 鏁版嵁搴撴湇鍔″櫒鐨勫湴鍧€
	Port         int    `json:"port" yaml:"port"`                     // 鏁版嵁搴撴湇鍔″櫒鐨勭鍙ｅ彿
	Config       string `json:"config" yaml:"config"`                 // 鏁版嵁搴撹繛鎺ョ殑閰嶇疆鍙傛暟锛屽椹卞姩銆佸瓧绗﹂泦绛?
	DBName       string `json:"db_name" yaml:"db_name"`               // 瑕佽繛鎺ョ殑鏁版嵁搴撳悕绉?
	Username     string `json:"username" yaml:"username"`             // 鐢ㄤ簬杩炴帴鏁版嵁搴撶殑鐢ㄦ埛鍚?
	Password     string `json:"password" yaml:"password"`             // 鐢ㄤ簬杩炴帴鏁版嵁搴撶殑瀵嗙爜
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"` // 鏈€澶х┖闂茶繛鎺ユ暟锛屾帶鍒惰繛鎺ユ睜涓殑绌洪棽杩炴帴鏁伴噺
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"` // 鏈€澶ф墦寮€杩炴帴鏁帮紝闄愬埗鍚屾椂鎵撳紑鐨勬暟鎹簱杩炴帴鏁伴噺
	LogMode      string `json:"log_mode" yaml:"log_mode"`             // 鏃ュ織妯″紡锛屼緥濡?"info" 鎴?"silent"锛岀敤浜庢帶鍒舵棩蹇楄緭鍑?
}

func (m Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DBName + "?" + m.Config
}
func (m Mysql) LogLevel() logger.LogLevel {
	switch strings.ToLower(m.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}