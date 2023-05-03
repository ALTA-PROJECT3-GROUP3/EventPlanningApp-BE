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

// DeleteHandler implements event.Handler
func (ev *eventController) DeleteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := helper.DecodeToken(c)
		if userId == 0 {
			c.Logger().Error("decode token is blank")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "jwt invalid", nil))
		}
		eventId, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			c.Logger().Error("Event tidak ditemukan")
			return errCnv
		}

		err := ev.service.DeleteBook(userId, uint(eventId))
		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan Delete Event", nil))
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "delete Event successfully", nil))
	}
}

// UpdateHandler implements event.Handler
func (ev *eventController) UpdateHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := helper.DecodeToken(c)
		if userId == 0 {
			c.Logger().Error("decode token is blank")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "jwt invalid", nil))
		}
		eventId, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			c.Logger().Error("Event tidak ditemukan")
			return errCnv
		}

		updateInput := EventRequest{}
		if err := c.Bind(&updateInput); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid input", nil))
		}

		file, _ := c.FormFile("pictures")

		updateEvent := event.Core{}
		copier.Copy(&updateEvent, &updateInput)
		err := ev.service.Update(userId, uint(eventId), updateEvent, file)
		if err != nil {
			c.Logger().Error("terjadi kesalahan Update Event", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan Update Event", nil))
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "update Event successfully", nil))
	}
}

// GetEventByIdHandler implements event.Handler
func (ev *eventController) GetEventByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		eventId, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			c.Logger().Error("terjadi kesalahan")
			return errCnv
		}
		data, err := ev.service.GetEventById(uint(eventId))
		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Failed, error read data", nil))
		}
		res := EventResponse{}
		copier.Copy(&res, &data)
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "detail book successfully displayed", res))
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
		data, err := ev.service.MyEvent(userId, pageNumber)
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
