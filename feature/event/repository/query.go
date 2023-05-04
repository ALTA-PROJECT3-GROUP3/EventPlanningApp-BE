package repository

import (
	"errors"

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

// DeleteBook implements event.Repository
func (ev *eventQuery) Delete(userId uint, id uint) error {
	tx := ev.db.Where("user_id = ?", userId).Delete(&Event{}, id)
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error")
		return errors.New("no data deleted")
	}
	if tx.Error != nil {
		log.Error("Event tidak ditemukan")
		return tx.Error
	}
	return nil
}

// Update implements event.Repository
func (ev *eventQuery) Update(userId uint, id uint, input event.Core) error {
	data := CoreToEvent(input)
	tx := ev.db.Model(&Event{}).Where("id = ? AND user_id = ?", id, userId).Updates(&data)
	if tx.RowsAffected < 1 {
		log.Error("Terjadi error saat Update Event")
		return errors.New("event no updated")
	}
	if tx.Error != nil {
		log.Error("Event tidak ditemukan")
		return tx.Error
	}
	return nil
}

// GetEventById implements event.Repository
func (ev *eventQuery) GetEventById(id uint) (event.Core, error) {
	tmp := Event{}
	tx := ev.db.Preload("Tickets").Preload("Comments").Where("id = ?", id).Select("events.id, events.name, events.host_name, events.description, events.date, events.location, events.is_paid, events.pictures").Joins("JOIN users ON events.user_id = users.id").Group("events.id").First(&tmp)
	if tx.Error != nil {
		log.Error("Event tidak ditemukan")
		return event.Core{}, tx.Error
	}

	return EventToCore(tmp), nil
}

// MyEvent implements event.Repository
func (ev *eventQuery) MyEvent(userId uint, limit int, offset int) ([]event.Core, error) {
	var eventsModel []Event
	tx := ev.db.Preload("Tickets").Preload("Comments").Limit(limit).Offset(offset).Where("user_id = ?", userId).Select("events.id, events.name, events.host_name, events.description, events.date, events.location, events.is_paid, events.pictures").Joins("JOIN users ON events.user_id = users.id").Group("events.id").Find(&eventsModel)
	if tx.Error != nil {
		log.Error("Terjadi error saat select Event")
		return nil, tx.Error
	}
	booksCoreAll := ListModelToCore(eventsModel)
	return booksCoreAll, nil
}

// SelectAll implements event.Repository
func (ev *eventQuery) SelectAll(limit int, offset int, name string) ([]event.Core, error) {
	nameSearch := "%" + name + "%"
	var eventsModel []Event
	tx := ev.db.Preload("Tickets").Preload("Comments").Limit(limit).Offset(offset).Where("events.name LIKE ?", nameSearch).Select("events.id, events.name, events.host_name, events.description, events.date, events.location, events.is_paid, events.pictures").Joins("JOIN users ON events.user_id = users.id").Group("events.id").Find(&eventsModel)
	if tx.Error != nil {
		log.Error("Terjadi error saat select Event")
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
