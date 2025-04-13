package database

import "server/global"

// ArticleLike 鏂囩珷鏀惰棌琛?
type ArticleLike struct {
	global.MODEL
	ArticleID string `json:"article_id"` // 鏂囩珷 ID
	UserID    uint   `json:"user_id"`    // 鐢ㄦ埛 ID
	User      User   `json:"-" gorm:"foreignKey:UserID"`
}
