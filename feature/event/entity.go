package event

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	Id          uint
	Name        string    `validate:"required"`
	HostName    string    `validate:"required"`
	Description string    `validate:"required"`
	Date        time.Time `validate:"required"`
	Location    string    `validate:"required"`
	IsPaid      bool
	Pictures    string `validate:"required"`
	UserID      uint
}

type Handler interface {
	AddHandler() echo.HandlerFunc
	GetAllHandler() echo.HandlerFunc
}

type UseCase interface {
	Add(newEvent Core, file *multipart.FileHeader) error
	GetAll(page int, name string) ([]Core, error)
}

type Repository interface {
	Insert(input Core) error
	SelectAll(limit, offset int, name string) ([]Core, error)
}
