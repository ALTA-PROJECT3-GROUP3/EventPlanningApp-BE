package repository

import (
	"time"

	cRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/repository"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(50);not null"`
	HostName    string    `gorm:"type:varchar(50);not null"`
	Description string    `gorm:"type:text;not null"`
	Date        time.Time `gorm:"type:datetime;not null"`
	Location    string    `gorm:"type:varchar(50);not null"`
	IsPaid      bool      `gorm:"default:false"`
	Pictures    string    `gorm:"type:text;not null"`
	UserID      uint
	Tickets     []Ticket
	Comments    []cRepo.Comment
}

type Ticket struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50);not null"`
	Quota   int    `gorm:"not null"`
	Price   int    `gorm:"not null"`
	EventID uint
}

func CoreToEvent(data event.Core) Event {
	return Event{
		Model:       gorm.Model{ID: data.Id},
		Name:        data.Name,
		HostName:    data.HostName,
		Description: data.Description,
		Date:        time.Time{},
		Location:    data.Location,
		IsPaid:      false,
		Pictures:    data.Pictures,
		UserID:      data.UserID,
	}
}
