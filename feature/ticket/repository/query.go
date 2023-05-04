package repository

import (
	"errors"

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

// Update implements ticket.Repository
func (tk *ticketModel) Update(userId uint, id uint, input ticket.Core) error {
	data := CoreToTicket(input)
	tx := tk.db.Model(&Ticket{}).Where("user_id = ? AND id = ?", userId, id).Updates(&data)
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error saat Update Ticket")
		return errors.New("ticket no updated")
	}
	if tx.Error != nil {
		log.Error("Ticket tidak ditemukan")
		return tx.Error
	}
	return nil
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
