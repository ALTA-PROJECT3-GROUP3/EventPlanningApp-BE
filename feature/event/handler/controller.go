package handler

import (
	"net/http"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type eventController struct {
	service event.UseCase
}

func New(us event.UseCase) event.Handler {
	return &eventController{
		service: us,
	}
}

// AddHandler implements event.Handler
func (ev *eventController) AddHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		addInput := EventRequest{}
		addInput.UserID = uint(helper.DecodeToken(c))
		if err := c.Bind(&addInput); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid input", nil))
		}

		newEvent := event.Core{}
		copier.Copy(&newEvent, &addInput)

		file, _ := c.FormFile("pictures")

		err := ev.service.Add(newEvent, file)
		if err != nil {
			c.Logger().Error("terjadi kesalahan saat add Event", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, err.Error(), nil))
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "add event successfully", nil))
	}
}
