package event

import (
	"mime/multipart"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket"
	"github.com/labstack/echo/v4"
)

type Core struct {
	Id          uint
	Name        string `validate:"required"`
	HostName    string `validate:"required"`
	Description string `validate:"required"`
	Date        string `validate:"required"`
	Location    string `validate:"required"`
	IsPaid      bool
	Pictures    string
	UserID      uint
	Tickets     []ticket.Core
}

// type Ticket struct {
// 	Id      uint
// 	Name    string
// 	Quota   int
// 	Price   int
// 	EventID uint
// }

type Handler interface {
	AddHandler() echo.HandlerFunc
	GetAllHandler() echo.HandlerFunc
	MyeventHandler() echo.HandlerFunc
	GetEventByIdHandler() echo.HandlerFunc
	UpdateHandler() echo.HandlerFunc
	DeleteHandler() echo.HandlerFunc
}

type UseCase interface {
	Add(newEvent Core, file *multipart.FileHeader) error
	GetAll(page int, name string) ([]Core, error)
	MyEvent(userId uint, page int) ([]Core, error)
	GetEventById(id uint) (Core, error)
	Update(userId uint, id uint, updateEvent Core, file *multipart.FileHeader) error
	DeleteBook(userId uint, id uint) error
}

type Repository interface {
	Insert(input Core) error
	SelectAll(limit, offset int, name string) ([]Core, error)
	MyEvent(userId uint, limit, offset int) ([]Core, error)
	GetEventById(id uint) (Core, error)
	Update(userId uint, id uint, input Core) error
	DeleteBook(userId uint, id uint) error
}
