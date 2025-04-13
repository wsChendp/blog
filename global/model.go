package global

import (
	"time"
	"gorm.io/gorm"
)

type MODEL struct{
	ID 		uint `json:"id" gorm:"primary_key;auto_increment;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;"`
}
