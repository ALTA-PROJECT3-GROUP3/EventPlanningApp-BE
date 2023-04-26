package handler

type commentRequest struct {
	EventID uint   `json:"event_id"`
	Comment string `json:"comment"`
}
