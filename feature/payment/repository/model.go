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

type ReservationsCore struct {
	PaymentID   uint
	UserID      uint
	EventID     uint
	OrderID     string
	PhoneNumber string
	PaymentType string
	Bank        string
	VA          string
	Status      string
	JoinDate    string
	Tickets     []Tickets
	GrandTotal  int
}

type Tickets struct {
	TicketID uint
	Name     string
	Quantity int
	Quota    int
	Price    int
}

func CoreToData(data payment.ReservationsCore) Payment {
	return Payment{
		Model:       gorm.Model{},
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
