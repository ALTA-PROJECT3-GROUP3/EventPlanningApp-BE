package main

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/config"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/database"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/routes"
	commentHandler "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/handler"
	commentRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/repository"
	commentLogic "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/usecase"
	uHandler "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/handler"
	uRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/repository"
	uLogic "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/usecase"
	"github.com/labstack/echo/v4"

	paymentHandler "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment/handler"
	paymentRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment/repository"
	paymentLogic "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment/usecase"

	eHandler "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event/handler"
	eRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event/repository"
	eLogic "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event/usecase"

	tHandler "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket/handler"
	tRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket/repository"
	tLogic "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket/usecase"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)
	database.Migrate(db)

	uMdl := uRepo.New(db)
	uSrv := uLogic.New(uMdl)
	uCtl := uHandler.New(uSrv)

	commentMdl := commentRepo.New(db)
	commentSrv := commentLogic.New(commentMdl)
	commentCtl := commentHandler.New(commentSrv)

	paymentMdl := paymentRepo.New(db)
	paymentSrv := paymentLogic.New(paymentMdl)
	PaymentCtl := paymentHandler.New(paymentSrv)

	eMdl := eRepo.New(db)
	eSrv := eLogic.New(eMdl)
	eCtl := eHandler.New(eSrv)

	tMdl := tRepo.New(db)
	tSrv := tLogic.New(tMdl)
	tCtl := tHandler.New(tSrv)

	routes.UserRoutes(e, uCtl)
	routes.CommentRoutes(e, commentCtl)
	routes.PaymentRoutes(e, PaymentCtl)
	routes.EventRoutes(e, eCtl)
	routes.TicketRoutes(e, tCtl)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("cannot start server", err.Error())
	}
}
