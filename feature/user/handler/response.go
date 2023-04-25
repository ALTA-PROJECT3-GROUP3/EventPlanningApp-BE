package handler

type GetUserByIdResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Pictures string `json:"pictures"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
