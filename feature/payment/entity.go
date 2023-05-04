package payment

import (
	"time"

	"github.com/labstack/echo/v4"
)

type PaymentCore struct {
	ID          uint
	PaymentType string
	Bank        string
	UserID      uint
	OrderID     string
	VA          string
	Status      string
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

type ReservationsCore struct {
	PaymentID   uint
	UserID      uint
	EventID     uint
	OrderID     string
	PhoneNumber string
	PaymentType string
	Bank        string
	VA          string
	Status      string
	JoinDate    string
	Tickets     []Tickets
	GrandTotal  int
}

type Tickets struct {
	TicketID uint
	Name     string
	Quantity int
	Quota    int
	Price    int
}

type Handler interface {
	CreateReservationHandler() echo.HandlerFunc
	MidtransNotification() echo.HandlerFunc
}

type UseCase interface {
	CreateReservationLogic(ReservationsCore) (ReservationsCore, error)
	UpdateTransaction(input PaymentCore) error
}

type Repository interface {
	CheckEvent(ReservationsCore) error
	CheckTicket(ReservationsCore) (ReservationsCore, error)
	InsertPayment(PaymentCore) error
	InsertPaymentDetails(PaymentDetailCore) error
	UpdateTransaction(input PaymentCore) error
}
