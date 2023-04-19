package database

import (
	cRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/repository"
	eRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event/repository"
	pRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment/repository"
	uRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&uRepo.User{})
	db.AutoMigrate(&eRepo.Event{})
	db.AutoMigrate(&eRepo.Ticket{})
	db.AutoMigrate(&pRepo.PaymentDetail{})
	db.AutoMigrate(&pRepo.Payment{})
	db.AutoMigrate(&cRepo.Comment{})
}
