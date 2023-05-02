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

	routes.UserRoutes(e, uCtl)
	routes.CommentRoutes(e, commentCtl)
	routes.PaymentRoutes(e, PaymentCtl)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal("cannot start server", err.Error())
	}
}
