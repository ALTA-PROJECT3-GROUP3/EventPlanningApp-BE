package handler

import (
	"net/http"
	"strings"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/echo/v4"
)

type ticketHandler struct {
	cln ticket.UseCase
}

func New(cln ticket.UseCase) ticket.Handler {
	return &ticketHandler{
		cln: cln,
	}
}

// CreateHandler implements ticket.Handler
func (tc *ticketHandler) CreateHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request = new(ticketRequest)
		userID := helper.DecodeToken(c)
		var newTicket ticket.Core
		if userID == 0 {
			c.Logger().Error("decode token is blank")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "jwt invalid", nil))
		}

		if err := c.Bind(&request); err != nil {
			c.Logger().Error("error on binding request create ticket")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid user indput", nil))
		}

		newTicket.Name = request.Name
		newTicket.Price = request.Price
		newTicket.Quota = request.Quota
		newTicket.EventID = request.EventID

		if err := tc.cln.Create(newTicket); err != nil {
			c.Logger().Error("error on calling CreateticketLogic")
			if strings.Contains(err.Error(), "connect") || strings.Contains(err.Error(), "table 'events' not found") || strings.Contains(err.Error(), "table 'tickets' not found") || strings.Contains(err.Error(), "server error") {
				c.Logger().Error("error on creating tickets, internal sever errors")
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
			}
			if strings.Contains(err.Error(), "bad request") {
				c.Logger().Error("bad request, event not found")
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "event is not exist or not has been deleted by owner", nil))
			}
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "succes add ticket", nil))
	}
}
