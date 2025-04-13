package database

// ArticleCategory 鏂囩珷绫诲埆琛?
type ArticleCategory struct {
	Category string `json:"category" gorm:"primaryKey"` // 绫诲埆
	Number   int    `json:"number"`                     // 缁熻鏁伴噺
}