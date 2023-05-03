package handler

type EventRequest struct {
	Name        string `json:"name" form:"name"`
	HostName    string `json:"host_name" form:"host_name"`
	Description string `json:"description" form:"description"`
	Date        string `json:"date" form:"date"`
	Location    string `json:"location" form:"location"`
	IsPaid      bool   `json:"is_paid" form:"is_paid"`
	Pictures    string `json:"pictures" form:"pictures"`
	UserID      uint   `json:"user_id" form:"user_id"`
	// TicketID    []int  `json:"ticket_id"form:"ticket_id"`
}

// type Ticket struct {
// 	Name    string `json:"name" form:"name"`
// 	Quota   int    `json:"quota" form:"quota"`
// 	Price   int    `json:"price" form:"price"`
// 	EventID uint   `json:"event_id" form:"event_id"`
// }
