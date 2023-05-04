package handler

import (
	cRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/repository"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event"
	tRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket/repository"
	"gorm.io/gorm"
)

type EventResponse struct {
	ID            uint            `json:"id"`
	Name          string          `json:"name"`
	HostName      string          `json:"host_name"`
	Description   string          `json:"description"`
	Date          string          `json:"date"`
	Location      string          `json:"location"`
	IsPaid        bool            `json:"is_paid"`
	AttendesQuota int             `json:"attendes_quota"`
	Pictures      string          `json:"pictures"`
	Tickets       []tRepo.Ticket  `json:"tickets"`
	Comments      []cRepo.Comment `json:"comments"`
}

func CoreToGetAllEventRespB(data event.Core) EventResponse {
	result := EventResponse{
		ID:            data.Id,
		Name:          data.Name,
		HostName:      data.HostName,
		Description:   data.Description,
		Date:          data.Date,
		Location:      data.Location,
		IsPaid:        false,
		AttendesQuota: data.AttendesQuota,
		Pictures:      data.Pictures,
	}

	for _, v := range data.Tickets {
		cticket := tRepo.Ticket{
			Model:   gorm.Model{ID: v.Id},
			Name:    v.Name,
			Quota:   v.Quota,
			Price:   v.Price,
			EventID: v.EventID,
			UserID:  v.UserID,
		}
		result.Tickets = append(result.Tickets, cticket)
	}

	for _, y := range data.Comments {
		cComment := cRepo.Comment{
			UserID:  y.UserID,
			EventID: y.EventID,
			Text:    y.Comment,
		}
		result.Comments = append(result.Comments, cComment)
	}
	return result
}

func CoreToGetAllEventResp(data []event.Core) []EventResponse {
	res := []EventResponse{}
	for _, val := range data {
		res = append(res, CoreToGetAllEventRespB(val))
	}
	return res
}
