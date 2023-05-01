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

// SelectAll implements event.Repository
func (ev *eventQuery) SelectAll(limit int, offset int, name string) ([]event.Core, error) {
	nameSearch := "%" + name + "%"
	var eventsModel []Event
	tx := ev.db.Limit(limit).Offset(offset).Where("events.name LIKE ?", nameSearch).Select("events.id, events.name, events.host_name, events.description, events.date, events.location, events.is_paid, events.pictures").Joins("JOIN users ON events.user_id = users.id").Group("events.id").Find(&eventsModel)
	if tx.Error != nil {
		log.Error("Terjadi error saat select Book")
		return nil, tx.Error
	}
	eventsCoreAll := ListModelToCore(eventsModel)
	return eventsCoreAll, nil
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
