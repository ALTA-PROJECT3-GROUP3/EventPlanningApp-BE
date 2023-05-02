package handler

import (
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
			reservation.Tickets[i].TicketID = req.Tickets[i].TicketID
			reservation.Tickets[i].Quantity = req.Tickets[i].Quantity
		}

		res, err := pc.service.CreateReservationLogic(reservation)
		if err != nil {
			c.Logger().Error("error on calling reservationlogic")
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		data.PaymentID = res.PaymentID
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "succes to join event", data))
	}
}
