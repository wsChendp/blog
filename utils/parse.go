package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseDuration 瑙ｆ瀽鎸佺画鏃堕棿瀛楃涓蹭负 time.Duration銆?
// 鎸佺画鏃堕棿瀛楃涓插簲鐢辨暟瀛楀€煎拰鏃堕棿鍗曚綅缁勬垚锛屽崟浣嶅彲浠ユ槸 "d" 琛ㄧず澶╋紝"h" 琛ㄧず灏忔椂锛?m" 琛ㄧず鍒嗛挓锛?s" 琛ㄧず绉掋€?
// 渚嬪锛?1d2h30m" 浼氳瑙ｆ瀽涓?1 澶┿€? 灏忔椂鍜?30 鍒嗛挓銆?
// 濡傛灉瀛楃涓蹭负绌烘垨鏍煎紡鏃犳晥锛屽垯杩斿洖閿欒銆?
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d) // 鍘婚櫎瀛楃涓蹭袱绔殑绌烘牸
	if len(d) == 0 {
		return 0, fmt.Errorf("empty duration string")
	}

	// 瀹氫箟姣忎釜鍗曚綅鍙婂叾瀵瑰簲鐨勬寔缁椂闂村€?
	unitPattern := map[string]time.Duration{
		"d": time.Hour * 24, // "d" 瀵瑰簲 24 灏忔椂
		"h": time.Hour,      // "h" 瀵瑰簲 1 灏忔椂
		"m": time.Minute,    // "m" 瀵瑰簲 1 鍒嗛挓
		"s": time.Second,    // "s" 瀵瑰簲 1 绉?
	}

	var totalDuration time.Duration // 鎬绘寔缁椂闂?
	// 閬嶅巻 "d"銆?h"銆?m"銆?s" 鍗曚綅
	for _, unit := range []string{"d", "h", "m", "s"} {
		// 鎻愬彇鎵€鏈変互褰撳墠鍗曚綅缁撳熬鐨勯儴鍒?
		for strings.Contains(d, unit) {
			// 鎵惧埌鍗曚綅鐨勪綅缃?
			unitIndex := strings.Index(d, unit)
			// 鎻愬彇鍗曚綅鍓嶉潰鐨勯儴鍒?
			part := d[:unitIndex]
			if part == "" {
				part = "0" // 濡傛灉閮ㄥ垎涓虹┖锛岄粯璁や负 0
			}
			// 灏嗛儴鍒嗚浆鎹负鏁存暟鍊?
			val, err := strconv.Atoi(part)
			if err != nil {
				return 0, fmt.Errorf("invalid duration part: %v", err)
			}
			// 灏嗚閮ㄥ垎鐨勬寔缁椂闂寸疮鍔犲埌鎬绘寔缁椂闂?
			totalDuration += time.Duration(val) * unitPattern[unit]
			// 浠庡瓧绗︿覆涓Щ闄ゅ凡澶勭悊鐨勯儴鍒?
			d = d[unitIndex+len(unit):]
		}
	}

	// 妫€鏌ユ槸鍚︽湁鍓╀綑鏈鐞嗙殑瀛楃
	if len(d) > 0 {
		return 0, fmt.Errorf("unrecognized duration format")
	}

	// 杩斿洖鎬荤殑鎸佺画鏃堕棿
	return totalDuration, nil
}