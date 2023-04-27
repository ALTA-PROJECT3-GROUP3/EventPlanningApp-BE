package handler

type ReservationRequest struct {
	EventID uint `json:"event_id"`
	Tickets []struct {
		TicketID uint `json:"ticket_id"`
		Quantity int  `json:"quantity"`
	} `json:"tickets"`
}
