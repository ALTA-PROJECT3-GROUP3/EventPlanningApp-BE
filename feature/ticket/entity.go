package ticket

import "github.com/labstack/echo/v4"

type Core struct {
	Id      uint
	Name    string
	Quota   int
	Price   int
	EventID uint
	UserID  uint
}

type Handler interface {
	CreateHandler() echo.HandlerFunc
}

type UseCase interface {
	Create(newTicket Core) error
}

type Repository interface {
	Insert(input Core) error
}
