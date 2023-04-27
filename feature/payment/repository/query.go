package repository

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment"
	"github.com/jinzhu/gorm"
)

type paymentModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.Repository {
	return &paymentModel{
		db: db,
	}
}