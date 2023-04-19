package repository

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID      uint
	GrandTotal  int       `gorm:"not null"`
	PaymentDate time.Time `gorm:"type:datetime"`
}

type PaymentDetail struct {
	ID        uint `gorm:"primary_key"`
	TicketID  uint
	PaymentID uint
	Qty       int `gorm:"not null"`
	Total     int `gorm:"not null"`
}
