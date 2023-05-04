package repository

import (
	"errors"
	"fmt"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type paymentModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.Repository {
	return &paymentModel{
		db: db,
	}
}

func (pm *paymentModel) CheckEvent(rsv payment.ReservationsCore) error {
	var count int64
	if err := pm.db.Raw("SELECT id FROM events WHERE id = ?", rsv.EventID).Count(&count).Error; err != nil {
		log.Error("error occured when finding event for reservation")
		return err
	}

	if count == 0 {
		log.Error("count is 0, event not found")
		return errors.New("event not found")
	}

	return nil
}

func (pm *paymentModel) CheckTicket(rsv payment.ReservationsCore) (payment.ReservationsCore, error) {
	var (
		count    int64
		ticketDB payment.Tickets
	)

	for _, ticket := range rsv.Tickets {
		if err := pm.db.Raw("SELECT id, name, quota, price, event_id FROM tickets WHERE id = ? AND event_id = ?", ticket.TicketID, rsv.EventID).Scan(&ticketDB).Count(&count).Error; err != nil {
			log.Error("error occurred in finding tickets for reservations")
			return rsv, err
		}
		fmt.Println(ticketDB)
		rsv.Tickets = append(rsv.Tickets, payment.Tickets{
			TicketID: ticketDB.TicketID,
			Name:     ticketDB.Name,
			Quantity: ticket.Quantity,
			Quota:    ticketDB.Quota,
			Price:    ticketDB.Price,
		})
	}
	fmt.Printf("rsv ticket %v", rsv.Tickets)
	return rsv, nil
}

func (pm *paymentModel) InsertPayment(paymentRsv payment.PaymentCore) error {
	var (
		paymentDB Payment
	)

	paymentDB.PaymentType = paymentRsv.PaymentType
	paymentDB.VA = paymentRsv.VA
	paymentDB.Bank = paymentRsv.Bank
	paymentDB.Status = paymentRsv.Status
	paymentDB.GrandTotal = paymentRsv.GrandTotal
	paymentDB.UserID = paymentRsv.UserID
	paymentDB.OrderID = paymentRsv.OrderID

	if err := pm.db.Create(&paymentDB).Error; err != nil {
		log.Error("error occured in inserting payment")
		return err
	}
	return nil

}

func (pm *paymentModel) InsertPaymentDetails(rsv payment.PaymentDetailCore) error {
	var (
		paymentDB PaymentDetail
	)

	paymentDB.TicketID = rsv.TicketID
	paymentDB.Total = rsv.Total
	paymentDB.PaymentID = rsv.PaymentID
	paymentDB.Qty = rsv.Qty

	if err := pm.db.Create(&paymentDB).Error; err != nil {
		log.Error("error occured in inserting payment detail")
		return err
	}
	return nil

}
