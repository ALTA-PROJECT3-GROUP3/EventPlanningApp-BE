package handler

type PaymentResponse struct {
	PaymentID    string `json:"payment_id"`
	InvcoiceDate string `json:"invoice_date"`
	Event_Name   string `json:"event:name"`
	Status       string `json:"status"`
	Total        int    `json:"total_price"`
	VA           string `json:"va_number"`
}
