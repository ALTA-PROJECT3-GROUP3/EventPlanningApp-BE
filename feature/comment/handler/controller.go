package handler

import (
	"net/http"
	"strings"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	cl comment.UseCase
}

func New(cl comment.UseCase) comment.Handler {
	return &commentHandler{
		cl: cl,
	}
}

func (ch *commentHandler) CreateCommentHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request = new(commentRequest)
		userID := helper.DecodeToken(c)
		var newComment comment.Core

		if userID == 0 {
			c.Logger().Error("decode token is blank")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "jwt invalid", nil))
		}

		if err := c.Bind(&request); err != nil {
			c.Logger().Error("error on binding request create comment")
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "invalid user indput", nil))
		}
		// return c.JSON(helper.ResponseFormat(http.StatusBadRequest, strconv.Itoa(int(request.EventID)), nil))
		newComment.UserID = userID
		newComment.EventID = request.EventID
		newComment.Comment = request.Comment

		if err := ch.cl.CreateCommentLogic(newComment); err != nil {
			c.Logger().Error("error on calling CreateCommentLogic")
			if strings.Contains(err.Error(), "connect") || strings.Contains(err.Error(), "table 'events' not found") || strings.Contains(err.Error(), "table 'comments' not found") || strings.Contains(err.Error(), "server error") {
				c.Logger().Error("error on creating comments, internal sever errors")
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
			}
			if strings.Contains(err.Error(), "bad request") {
				c.Logger().Error("bad request, event not found")
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "event is not exist or not has been deleted by owner", nil))
			}
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "succes add comment", nil))
	}
}
