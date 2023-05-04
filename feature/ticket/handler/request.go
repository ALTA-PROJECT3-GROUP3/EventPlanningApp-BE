package handler

type ticketRequest struct {
	Name    string `json:"name"`
	Quota   int    `json:"quota"`
	Price   int    `json:"price"`
	EventID uint   `json:"event_id"`
}

type updateRequest struct {
	Name  string `json:"name"`
	Quota int    `json:"quota"`
	Price int    `json:"price"`
}
