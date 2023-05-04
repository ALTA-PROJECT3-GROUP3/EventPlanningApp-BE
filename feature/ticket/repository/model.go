package repository

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50);not null"`
	Quota   int    `gorm:"not null"`
	Price   int    `gorm:"not null"`
	EventID uint
	UserID  uint
}

func CoreToTicket(data ticket.Core) Ticket {
	return Ticket{
		Model:   gorm.Model{ID: data.Id},
		Name:    data.Name,
		Quota:   data.Quota,
		Price:   data.Price,
		EventID: data.EventID,
		UserID:  data.UserID,
	}
}
