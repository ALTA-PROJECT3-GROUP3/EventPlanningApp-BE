package handler

import (
	"net/http"
	"strconv"
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

func (uc *userController) DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := helper.DecodeToken(c)
		if userId == 0 {
			c.Logger().Error("decode token is blank")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "jwt invalid", nil))
		}

		userPath, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error("cannot use path param", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusNotFound, "path invalid", nil))
		}

		if userId != uint(userPath) {
			c.Logger().Error("userpath is not equal with userId")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "user are not authorized to delete other user account", nil))
		}

		if err = uc.service.DeleteUserLogic(uint(userPath)); err != nil {
			c.Logger().Error("error in calling DeletUserLogic")
			if strings.Contains(err.Error(), "user not found") {
				c.Logger().Error("error in calling DeletUserLogic, user not found")
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "user not found", nil))

			} else if strings.Contains(err.Error(), "cannot delete") {
				c.Logger().Error("error in calling DeletUserLogic, cannot delete")
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "server error in delete user", nil))
			}

			c.Logger().Error("error in calling DeletUserLogic, cannot delete")
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "server error in delete user", nil))

		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "succes to delete user", nil))
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

func (uc *userController) UserProfileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data = new(GetUserByIdResponsestruct)
		userId := helper.DecodeToken(c)
		if userId == 0 {
			c.Logger().Error("decode token is blank")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "jwt invalid", nil))
		}

		userPath, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error("cannot use path param", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusNotFound, "path invalid", nil))
		}

		result, err := uc.service.UserProfileLogic(uint(userPath))
		if err != nil {
			c.Logger().Error("error on calling userpofilelogic")
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		data.Email = result.Email
		data.Name = result.Name
		data.Pictures = result.Picture

		return c.JSON(helper.ResponseFormat(http.StatusOK, "succes to check user profile", data))
	}
}
