package database

import (
	"server/global"
	"server/model/appTypes"

	"github.com/google/uuid"
)

type User struct {
	global.MODEL
	UUID      uuid.UUID 		`json:"uuid" gorm:"type:char(36);unique;"`
	Username  string 			`json:"username" gorm:"type:varchar(100);unique;"`
	Password  string 			`json:"-"`
	Email     string 			`json:"email"`
	Openid    string 			`json:"openid"`
	Avatar    string 			`json:"avatar"`
	Address   string 			`json:"address"`
	Signature string 			`json:"signature" gorm:"defualt:'签名是空白的，这个用户似乎比较低调'"`
	RoleID    appTypes.RoleID 	`json:"role_id"`
	Register  appTypes.Register `json:"register"`
	Freeze    bool 				`json:"freeze"`
}

// gorm
func (User) TableName() string {
	return "user"
}
