package config

// Qiniu 涓冪墰浜戦厤缃紝璇︽儏璇疯 https://www.qiniu.com/
type Qiniu struct {
	Zone          string `json:"zone" yaml:"zone"`                       // 瀛樺偍鍖哄煙
	Bucket        string `json:"bucket" yaml:"bucket"`                   // 绌洪棿鍚嶇О
	ImgPath       string `json:"img_path" yaml:"img_path"`               // CDN 鍔犻€熷煙鍚?
	AccessKey     string `json:"access_key" yaml:"access_key"`           // 绉橀挜 AK
	SecretKey     string `json:"secret_key" yaml:"secret_key"`           // 绉橀挜 SK
	UseHTTPS      bool   `json:"use_https" yaml:"use_https"`             // 鏄惁浣跨敤 https
	UseCdnDomains bool   `json:"use_cdn_domains" yaml:"use_cdn_domains"` // 涓婁紶鏄惁浣跨敤 CDN 涓婁紶鍔犻€?
}