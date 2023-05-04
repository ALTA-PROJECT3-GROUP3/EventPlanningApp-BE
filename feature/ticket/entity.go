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
	UpdateHandler() echo.HandlerFunc
}

type UseCase interface {
	Create(newTicket Core) error
	Update(userId uint, id uint, updateTicket Core) error
}

type Repository interface {
	Insert(input Core) error
	Update(userId uint, id uint, input Core) error
}
