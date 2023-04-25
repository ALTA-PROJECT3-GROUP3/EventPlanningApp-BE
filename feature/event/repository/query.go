package repository

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type eventQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) event.Repository {
	return &eventQuery{
		db: db,
	}
}

// Insert implements event.Repository
func (ev *eventQuery) Insert(input event.Core) error {
	data := CoreToEvent(input)
	tx := ev.db.Create(&data)
	if tx.Error != nil {
		log.Error("Terjadi error saat Insert Event")
		return tx.Error
	}
	return nil
}
