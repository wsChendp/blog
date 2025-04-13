package config

// Website 缃戠珯淇℃伅
type Website struct {
	Logo                 string `json:"logo" yaml:"logo"`
	FullLogo             string `json:"full_logo" yaml:"full_logo"`
	Title                string `json:"title" yaml:"title"`                                   // 缃戠珯鏍囬
	Slogan               string `json:"slogan" yaml:"slogan"`                                 // 缃戠珯鏍囪
	SloganEn             string `json:"slogan_en" yaml:"slogan_en"`                           // 鑻辨枃鏍囪
	Description          string `json:"description" yaml:"description"`                       // 缃戠珯鎻忚堪
	Version              string `json:"version" yaml:"version"`                               // 缃戠珯鐗堟湰
	CreatedAt            string `json:"created_at" yaml:"created_at"`                         // 鍒涘缓鏃堕棿
	IcpFiling            string `json:"icp_filing" yaml:"icp_filing"`                         // ICP 澶囨
	PublicSecurityFiling string `json:"public_security_filing" yaml:"public_security_filing"` // 鍏畨澶囨
	BilibiliURL          string `json:"bilibili_url" yaml:"bilibili_url"`                     // Bilibili 閾炬帴
	GiteeURL             string `json:"gitee_url" yaml:"gitee_url"`                           // Gitee 閾炬帴
	GithubURL            string `json:"github_url" yaml:"github_url"`                         // GitHub 閾炬帴
	Name                 string `json:"name" yaml:"name"`                                     // 鏄电О
	Job                  string `json:"job" yaml:"job"`                                       // 鑱屼笟
	Address              string `json:"address" yaml:"address"`                               // 鍦板潃
	Email                string `json:"email" yaml:"email"`                                   // 閭
	QQImage              string `json:"qq_image" yaml:"qq_image"`                             // QQ 鍥剧墖閾炬帴
	WechatImage          string `json:"wechat_image" yaml:"wechat_image"`                     // 寰俊鍥剧墖閾炬帴
}