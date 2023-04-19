package repository

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint
	EventID uint
	Text    string `gorm:"type:text"`
}
