package handler

import (
	cRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/repository"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event"
	tRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket/repository"
)

type EventResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	HostName    string `json:"host_name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Location    string `json:"location"`
	IsPaid      bool   `json:"is_paid"`
	Pictures    string `json:"pictures"`
	// UserID      uint           `json:"user_id"`
	Tickets  []tRepo.Ticket `json:"tickets"`
	Comments []cRepo.Comment
}

func CoreToGetAllEventRespB(data event.Core) EventResponse {
	result := EventResponse{
		ID:          data.Id,
		Name:        data.Name,
		HostName:    data.HostName,
		Description: data.Description,
		Date:        data.Date,
		Location:    data.Location,
		IsPaid:      false,
		Pictures:    data.Pictures,
		// UserID:      data.UserID,
	}

	for _, v := range data.Tickets {
		cticket := tRepo.Ticket{
			Name:    v.Name,
			Quota:   v.Quota,
			Price:   v.Price,
			EventID: v.EventID,
		}
		result.Tickets = append(result.Tickets, cticket)
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
