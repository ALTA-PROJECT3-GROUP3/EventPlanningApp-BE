package routes

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoutes(e *echo.Echo, uc user.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/login", uc.LoginHandler())
	e.POST("/users", uc.RegisterHandler())
	e.GET("/users/:id", uc.UserProfileHandler(), helper.JWTMiddleware())
	e.DELETE("/users/:id", uc.DeleteUserHandler(), helper.JWTMiddleware())
}
