package handler

import (
	"time"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event"
)

type EventResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	HostName    string    `json:"host_name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
	IsPaid      bool      `json:"is_paid"`
	Pictures    string    `json:"pictures"`
	UserID      uint      `json:"user_id"`
}

func CoreToGetAllEventRespB(data event.Core) EventResponse {
	return EventResponse{
		ID:          data.Id,
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

func CoreToGetAllEventResp(data []event.Core) []EventResponse {
	res := []EventResponse{}
	for _, val := range data {
		res = append(res, CoreToGetAllEventRespB(val))
	}
	return res
}
