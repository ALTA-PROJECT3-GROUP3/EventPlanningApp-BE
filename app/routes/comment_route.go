package routes

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CommentRoutes(e *echo.Echo, ch comment.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.POST("/comments", ch.CreateCommentHandler(), helper.JWTMiddleware())
}
