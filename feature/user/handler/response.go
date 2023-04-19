package handler

type GetUserByIdResponsestruct struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Pictures string `json:"pictrures"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
