package routes

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TicketRoutes(e *echo.Echo, tc ticket.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.POST("/tickets", tc.CreateHandler(), helper.JWTMiddleware())
	e.PUT("/tickets", tc.CreateHandler(), helper.JWTMiddleware())
}
