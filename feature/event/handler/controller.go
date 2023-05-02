package handler

import (
	"net/http"
	"strconv"

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

// MyeventHandler implements event.Handler
func (ev *eventController) MyeventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := helper.DecodeToken(c)
		if userId == 0 {
			c.Logger().Error("decode token is blank")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "jwt invalid", nil))
		}

		var pageNumber int = 1
		pageParam := c.QueryParam("page")
		if pageParam != "" {
			pageConv, errConv := strconv.Atoi(pageParam)
			if errConv != nil {
				c.Logger().Error("terjadi kesalahan")
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Failed, page must number", nil))
			} else {
				pageNumber = pageConv
			}
		}

		// nameParam := c.QueryParam("name")
		data, err := ev.service.MyEvent(int(userId), pageNumber)
		if err != nil {
			c.Logger().Error("terjadi kesalahan")
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Failed, error read data", nil))
		}
		dataResponse := CoreToGetAllEventResp(data)
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "get all event successfully", dataResponse))
	}
}

// GetAllHandler implements event.Handler
func (ev *eventController) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var pageNumber int = 1
		pageParam := c.QueryParam("page")
		if pageParam != "" {
			pageConv, errConv := strconv.Atoi(pageParam)
			if errConv != nil {
				c.Logger().Error("terjadi kesalahan")
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Failed, page must number", nil))
			} else {
				pageNumber = pageConv
			}
		}

		nameParam := c.QueryParam("name")
		data, err := ev.service.GetAll(pageNumber, nameParam)
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Failed, error read data", nil))
		}
		dataResponse := CoreToGetAllEventResp(data)
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "get all event successfully", dataResponse))
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
