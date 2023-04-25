package event

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	Id          uint
	Name        string
	HostName    string
	Description string
	Date        time.Time
	Location    string
	IsPaid      bool
	Pictures    string
	UserID      uint
}

type Handler interface {
	AddHandler() echo.HandlerFunc
}

type UseCase interface {
	Add(newEvent Core, file *multipart.FileHeader) error
}

type Repository interface {
	Insert(input Core) error
}
