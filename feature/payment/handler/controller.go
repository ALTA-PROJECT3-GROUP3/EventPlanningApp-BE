package handler

import (
	"fmt"
	"net/http"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type paymentController struct {
	service payment.UseCase
}

func New(ps payment.UseCase) payment.Handler {
	return &paymentController{
		service: ps,
	}
}

// MidtransNotification implements payment.Handler
func (pc *paymentController) MidtransNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := CheckTransactionRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		errUpdate := pc.service.UpdateTransaction(checkTransactionRequestToCore(input))
		if errUpdate != nil {
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "error update", nil))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get notification from midtrans",
		})
	}
}

func (pc *paymentController) CreateReservationHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(ReservationRequest)
		data := new(PaymentResponse)
		if err := c.Bind(&req); err != nil {
			log.Error("error on binding reservation request")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "bad request", nil))
		}

		userId := helper.DecodeToken(c)
		if userId == 0 {
			c.Logger().Error("decode token is blank")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "jwt invalid", nil))
		}
		var reservation payment.ReservationsCore
		reservation.UserID = userId
		reservation.EventID = req.EventID
		reservation.PaymentType = req.PaymentMethod

		for i := 0; i < len(req.Tickets); i++ {
			fmt.Println(req.Tickets)
			fmt.Println(len(reservation.Tickets))
			tmp := payment.Tickets{
				TicketID: req.Tickets[i].TicketID,
				Quantity: req.Tickets[i].Quantity,
			}
			reservation.Tickets = append(reservation.Tickets, tmp)
		}

		res, err := pc.service.CreateReservationLogic(reservation)
		if err != nil {
			c.Logger().Error("error on calling reservationlogic")
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		data.Event_Name = res.Tickets[0].Name
		data.PaymentID = res.OrderID
		data.Status = res.Status
		data.VA = res.VA
		data.Total = res.GrandTotal
		data.InvcoiceDate = res.JoinDate
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "succes to join event", data))
	}
}
