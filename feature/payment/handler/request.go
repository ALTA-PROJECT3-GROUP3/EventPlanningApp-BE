package handler

import "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment"

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

type CheckTransactionRequest struct {
	OrderID string `json:"order_id" form:"order_id"`
	Status  string `json:"status" form:"status"`
}

func checkTransactionRequestToCore(data CheckTransactionRequest) payment.PaymentCore {
	return payment.PaymentCore{
		OrderID: data.OrderID,
		Status:  data.Status,
	}
}
