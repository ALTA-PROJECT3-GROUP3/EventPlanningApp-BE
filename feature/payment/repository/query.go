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

// UpdateTransaction implements payment.Repository
func (pm *paymentModel) UpdateTransaction(input payment.PaymentCore) error {
	cnv := CoreToData(input)

	err := pm.db.Where("order_id = ?", cnv.OrderID).Updates(&cnv)
	if err != nil {
		log.Error("query error", err.Error)
		return errors.New("server error")
	}
	return nil
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

	for i := 0; i < len(rsv.Tickets); i++ {
		if err := pm.db.Raw("SELECT id, name, quota, price, event_id FROM tickets WHERE id = ? AND event_id = ?", rsv.Tickets[i].TicketID, rsv.EventID).Scan(&ticketDB).Count(&count).Error; err != nil {
			log.Error("error occurred in finding tickets for reservations")
			return rsv, err
		}
		rsv.Tickets[i].Name = ticketDB.Name
		rsv.Tickets[i].Quota = ticketDB.Quota
		rsv.Tickets[i].Price = ticketDB.Price
	}
	fmt.Printf("rsv Name: %v\n", rsv.Tickets[0].Name)
	fmt.Printf("rsv Quantity: %v\n", rsv.Tickets[0].Quantity)
	fmt.Printf("rsv Quota: %v\n", rsv.Tickets[0].Quota)
	fmt.Printf("rsv Price: %v\n", rsv.Tickets[0].Price)
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
