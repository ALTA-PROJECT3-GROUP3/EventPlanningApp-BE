package handler

type GetUserByIdResponsestruct struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Pictures string `json:"pictrures"`
}
