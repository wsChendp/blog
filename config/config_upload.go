package config

type Upload struct {
    Size int    `json:"size" yaml:"size"` // 鍥剧墖涓婁紶鐨勫ぇ灏忥紝鍗曚綅 MB
	Path string `json:"path" yaml:"path"` // 鍥剧墖涓婁紶鐨勭洰褰?
}