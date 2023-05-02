package routes

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func PaymentRoutes(e *echo.Echo, pc payment.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/reservations", pc.CreateReservationHandler(), helper.JWTMiddleware())
}
