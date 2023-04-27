package payment

import (
	"time"

	"github.com/labstack/echo/v4"
)

type PaymentCore struct {
	ID          uint
	UserID      uint
	GrandTotal  int
	PaymentDate time.Time
}

type PaymentDetailCore struct {
	ID        uint
	TicketID  uint
	PaymentID uint
	Qty       int
	Total     int
}

type Handler interface {
	CreateReservationHandler() echo.HandlerFunc
}

type UseCase interface {
	CreateReservationLogic() error
}

type Repository interface {
	InsertReservation() error
}
