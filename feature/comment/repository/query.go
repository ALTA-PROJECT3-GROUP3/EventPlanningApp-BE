package repository

import (
	"errors"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type commentModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.Repository {
	return &commentModel{
		db: db,
	}
}

func (cm *commentModel) InsertComment(newComment comment.Core) error {
	var insert Comment
	insert.UserID = newComment.UserID
	insert.EventID = newComment.EventID
	insert.Text = newComment.Comment

	var count int64
	err := cm.db.Table("events").Where("id = ?", insert.EventID).Count(&count).Error
	if err != nil {
		log.Error("error occurs at find events in comment query")
		return err
	}

	if count == 0 {
		log.Error("count is zero, event id is not exist")
		return errors.New("event not found")
	}
	
	if err := cm.db.Table("comments").Create(insert).Error; err != nil {
		log.Error("error occurs at create comments in comment query")
		return err
	}

	return nil
}
