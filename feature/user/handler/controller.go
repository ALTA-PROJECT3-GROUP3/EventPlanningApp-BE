package handler

import (
	"net/http"
	"strings"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/echo/v4"
)

type userController struct {
	service user.UseCase
}

func New(us user.UseCase) user.Handler {
	return &userController{
		service: us,
	}
}

func (uc *userController) RegisterHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterInput{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("error on bind register input", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid input", nil))
		}

		err := uc.service.RegisterUser(user.Core{
			Name:     input.Name,
			Email:    input.Email,
			Username: input.Username,
			Password: input.Password,
		})
		if err != nil {
			c.Logger().Error("error on calling userLogic", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "succes to create user", nil))
	}
}

func (uc *userController) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("error on bind login input", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid input", nil))
		}

		res, err := uc.service.LogInLogic(input.Username, input.Password)
		if err != nil {
			c.Logger().Error("error on calling Login Logic", err.Error())

			if strings.Contains(err.Error(), "not exist") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "username does not exist, please sign up", nil))
			} else if strings.Contains(err.Error(), "wrong") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "password is wrong please try again", nil))
			} else if strings.Contains(err.Error(), "blank") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "username is blank please try again", nil))
			}
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "internal server error", nil))
		}
		token, err := helper.GenerateToken(res.ID)
		if err != nil {
			c.Logger().Error("error on generation token", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil))
		}
		var data = new(LoginResponse)
		data.ID = int(res.ID)
		data.Name = res.Name
		data.Username = res.Username
		data.Token = token

		return c.JSON(helper.ResponseFormat(http.StatusOK, "succes login!", data))
	}
}
