package handler

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Pictures string `json:"pictures"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
