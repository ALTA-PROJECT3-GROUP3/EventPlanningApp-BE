package repository

import (
	"time"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PaymentType string
	Bank        string
	OrderID     string
	VA          string
	Status      string
	GrandTotal  int       `gorm:"not null"`
	PaymentDate time.Time `gorm:"type:datetime;default:null"`
	UserID      uint
}

type PaymentDetail struct {
	ID        uint `gorm:"primary_key"`
	TicketID  uint
	PaymentID uint
	Qty       int `gorm:"not null"`
	Total     int `gorm:"not null"`
}

func CoreToData(data payment.PaymentCore) Payment {
	return Payment{
		Model:       gorm.Model{ID: data.ID},
		PaymentType: data.PaymentType,
		Bank:        data.Bank,
		OrderID:     data.OrderID,
		VA:          data.VA,
		Status:      data.Status,
		GrandTotal:  data.GrandTotal,
		PaymentDate: time.Time{},
		UserID:      data.UserID,
	}
}
