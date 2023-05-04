package repository

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment"
	cRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/repository"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket"
	tRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket/repository"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);not null"`
	HostName    string `gorm:"type:varchar(50);not null"`
	Description string `gorm:"type:text;not null"`
	Date        string `gorm:"type:varchar(50);not null"`
	Location    string `gorm:"type:varchar(50);not null"`
	IsPaid      bool   `gorm:"default:false"`
	Pictures    string `gorm:"type:text"`
	UserID      uint
	Tickets     []tRepo.Ticket
	Comments    []cRepo.Comment
}

func CoreToEvent(data event.Core) Event {
	return Event{
		Model:       gorm.Model{ID: data.Id},
		Name:        data.Name,
		HostName:    data.HostName,
		Description: data.Description,
		Date:        data.Date,
		Location:    data.Location,
		IsPaid:      false,
		Pictures:    data.Pictures,
		UserID:      data.UserID,
	}
}

func EventToCore(data Event) event.Core {
	result := event.Core{
		Id:          data.ID,
		Name:        data.Name,
		HostName:    data.HostName,
		Description: data.Description,
		Date:        data.Date,
		Location:    data.Location,
		IsPaid:      false,
		Pictures:    data.Pictures,
		UserID:      data.UserID,
	}

	for _, v := range data.Tickets {
		cticket := ticket.Core{
			Id:      v.ID,
			Name:    v.Name,
			Quota:   v.Quota,
			Price:   v.Price,
			EventID: v.EventID,
		}
		result.Tickets = append(result.Tickets, cticket)
	}

	for _, y := range data.Comments {
		cComment := comment.Core{
			UserID:  y.UserID,
			EventID: y.EventID,
			Comment: y.Text,
		}
		result.Comments = append(result.Comments, cComment)
	}
	return result
}

func ListModelToCore(dataModel []Event) []event.Core {
	var dataCore []event.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, EventToCore(v))
	}
	return dataCore
}
