package config

// Zap 鏃ュ織閰嶇疆锛岃鎯呭彲鍙傝€冧竷绫崇殑鍗氬 https://liwenzhou.com/posts/Go/zap/
type Zap struct {
	Level          string `json:"level" yaml:"level"`                       // 鏃ュ織绛夌骇锛屾棤鐗规畩闇€姹傦紝鐢?info 鍗冲彲
	Filename       string `json:"filename" yaml:"filename"`                 // 鏃ュ織鏂囦欢鐨勪綅缃?
	MaxSize        int    `json:"max_size" yaml:"max_size"`                 // 鍦ㄨ繘琛屽垏鍓蹭箣鍓嶏紝鏃ュ織鏂囦欢鐨勬渶澶уぇ灏忥紙浠B涓哄崟浣嶏級
	MaxBackups     int    `json:"max_backups" yaml:"max_backups"`           // 淇濈暀鏃ф枃浠剁殑鏈€澶т釜鏁?
	MaxAge         int    `json:"max_age" yaml:"max_age"`                   // 淇濈暀鏃ф枃浠剁殑鏈€澶уぉ鏁?
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"` // 鏄惁鍦ㄦ帶鍒跺彴鎵撳嵃鏃ュ織锛宼rue 琛ㄧず鎵撳嵃锛宖alse 琛ㄧず涓嶆墦鍗?
}