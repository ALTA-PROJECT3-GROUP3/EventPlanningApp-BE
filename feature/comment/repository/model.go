package repository

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint
	EventID uint
	Text    string
}

func CoreToComment(data comment.Core) Comment {
	return Comment{
		UserID:  data.UserID,
		EventID: data.EventID,
		Text:    data.Comment,
	}
}
