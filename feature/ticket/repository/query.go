package repository

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ticketModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) ticket.Repository {
	return &ticketModel{
		db: db,
	}
}

// Insert implements ticket.Repository
func (tc *ticketModel) Insert(input ticket.Core) error {
	data := CoreToTicket(input)
	tx := tc.db.Create(&data)
	if tx.Error != nil {
		log.Error("Terjadi error saat create")
		return tx.Error
	}
	return nil

}
