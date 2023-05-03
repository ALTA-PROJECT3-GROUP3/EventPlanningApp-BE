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
	// var count int64
	// err := tc.db.Table("events").Where("id = ?", input.EventID).Count(&count).Error
	// if err != nil {
	// 	log.Error("error occurs at find events in ticket query")
	// 	return err
	// }

	// if count == 0 {
	// 	log.Error("count is zero, event id is not exist")
	// 	return errors.New("event not found")
	// }

	// if err := tc.db.Table("tickets").Create(input).Error; err != nil {
	// 	log.Error("error occurs at create tickets in event query")
	// 	return err
	// }

	// return nil
}

// package repository

// import (
// 	"errors"

// 	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket"
// 	"github.com/labstack/gommon/log"
// 	"gorm.io/gorm"
// )

// type ticketModel struct {
// 	db *gorm.DB
// }

// func New(db *gorm.DB) ticket.Repository {
// 	return &ticketModel{
// 		db: db,
// 	}
// }

// // Insert implements ticket.Repository
// func (tc *ticketModel) Insert(input ticket.Core) error {
// 	var count int64
// 	err := tc.db.Table("events").Where("id = ?", input.EventID).Count(&count).Error
// 	if err != nil {
// 		log.Error("error occurs at find events in ticket query")
// 		return err
// 	}

// 	if count == 0 {
// 		log.Error("count is zero, event id is not exist")
// 		return errors.New("event not found")
// 	}

// 	if err := tc.db.Table("tickets").Create(input).Error; err != nil {
// 		log.Error("error occurs at create tickets in comment query")
// 		return err
// 	}

// 	return nil

// }
