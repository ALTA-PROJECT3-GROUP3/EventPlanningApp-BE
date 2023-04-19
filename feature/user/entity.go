package user

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	Name     string
	Email    string
	Password string
	Username string
	Picture  string
}

type Handler interface {
	RegisterHandler() echo.HandlerFunc
}

type UseCase interface {
	RegisterUser(newUser Core) error
}

type Repository interface {
	InsertUser(newUser Core) error
}
