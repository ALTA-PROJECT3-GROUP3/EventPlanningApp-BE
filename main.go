package main

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/config"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/database"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/routes"
	uHandler "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/handler"
	uRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/repository"
	uLogic "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)
	database.Migrate(db)

	uMdl := uRepo.New(db)
	uSrv := uLogic.New(uMdl)
	uCtl := uHandler.New(uSrv)

	routes.UserRoutes(e, uCtl)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("cannot start server", err.Error())
	}
}
