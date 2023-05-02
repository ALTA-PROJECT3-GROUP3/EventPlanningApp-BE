package handler

type ReservationRequest struct {
	EventID       uint              `json:"event_id"`
	PhoneNumber   string            `json:"phone_number"`
	PaymentMethod string            `json:"payment_method"`
	Bank          string            `json:"bank"`
	Tickets       []ReservationItem `json:"tickets"`
}

type ReservationItem struct {
	TicketID uint `json:"ticket_id"`
	Quantity int  `json:"quantity"`
}
