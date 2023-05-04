package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/config"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment"
	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type paymentLogic struct {
	pm payment.Repository
}

func New(pr payment.Repository) payment.UseCase {
	return &paymentLogic{
		pm: pr,
	}
}

func (pl *paymentLogic) CreateReservationLogic(rsv payment.ReservationsCore) (payment.ReservationsCore, error) {
	var (
		pay       payment.PaymentCore
		payDetail payment.PaymentDetailCore
	)

	if err := pl.pm.CheckEvent(rsv); err != nil {
		log.Error("error in check event")
		return payment.ReservationsCore{}, err
	}

	reservation, err := pl.pm.CheckTicket(rsv)
	if err != nil {
		log.Error("error in check ticket")
		return payment.ReservationsCore{}, err
	}

	for _, ticket := range reservation.Tickets {
		fmt.Println(reservation.Tickets)
		if ticket.Quota < ticket.Quantity {
			log.Printf("ticket quota is %d, and quantity to buy is %d ticket are out of quota", ticket.Quota, ticket.Quantity)
			return payment.ReservationsCore{}, errors.New("out of quota")
		}
		reservation.GrandTotal += (ticket.Quota * ticket.Quantity)
	}

	reservationCharge, err := ChargePayment(reservation)
	if err != nil {
		log.Error("error occured when chargin to midtrans")
		return payment.ReservationsCore{}, nil
	}

	pay.GrandTotal = reservationCharge.GrandTotal
	pay.PaymentType = reservationCharge.PaymentType
	pay.Status = reservationCharge.Status
	pay.UserID = reservation.UserID
	pay.OrderID = reservationCharge.OrderID
	pay.VA = reservationCharge.VA
	pay.Bank = reservationCharge.Bank

	if err := pl.pm.InsertPayment(pay); err != nil {
		log.Error("error occured in inserting payment")
		return payment.ReservationsCore{}, err
	}

	for _, ticket := range reservationCharge.Tickets {
		payDetail.TicketID = ticket.TicketID
		payDetail.PaymentID = reservation.PaymentID
		payDetail.Qty = ticket.Quantity
		payDetail.Total = ticket.Quantity * ticket.Price
		if err := pl.pm.InsertPaymentDetails(payDetail); err != nil {
			log.Error("error occured when inserting payment details")
			return payment.ReservationsCore{}, err
		}
	}

	return reservationCharge, nil
}

func ChargePayment(rsv payment.ReservationsCore) (payment.ReservationsCore, error) {

	var (
		tmp = make([]midtrans.ItemDetails, len(rsv.Tickets))
		c   coreapi.Client
	)
	c.New(config.MidtransServerKey, midtrans.Sandbox)

	for i := 0; i < len(rsv.Tickets); i++ {
		tmp[i].Name = rsv.Tickets[i].Name
		tmp[i].Qty = int32(rsv.Tickets[i].Quantity)
		tmp[i].Price = int64(rsv.Tickets[i].Price)
	}

	reqTransactions := &coreapi.ChargeReq{
		PaymentType: coreapi.CoreapiPaymentType(rsv.PaymentType),
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(time.Now().Unix())),
			GrossAmt: int64(rsv.GrandTotal),
		},
		Items: &tmp,
	}

	resp, err := c.ChargeTransaction(reqTransactions)
	if err != nil {
		log.Error("error occured in charging the transaction to midtrans")
		return payment.ReservationsCore{}, err
	}
	rsv.VA = resp.VaNumbers[1].VANumber
	rsv.Status = resp.TransactionStatus
	rsv.OrderID = resp.OrderID

	return rsv, nil
}
