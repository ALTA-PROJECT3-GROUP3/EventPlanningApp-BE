package handler

import "time"

type EventRequest struct {
	Name        string    `json:"name" form:"name"`
	HostName    string    `json:"host_name" form:"host_name"`
	Description string    `json:"description" form:"description"`
	Date        time.Time `json:"date" form:"date"`
	Location    string    `json:"location" form:"location"`
	IsPaid      bool      `json:"is_paid" form:"is_paid"`
	Pictures    string    `json:"pictures" form:"pictures"`
	UserID      uint      `json:"user_id" form:"user_id"`
}
