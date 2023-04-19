package repository

import (
	cRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment/repository"
	eRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event/repository"
	pRepo "github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/payment/repository"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);not null"`
	Username string `gorm:"type:varchar(50);unique;not null"`
	Email    string `gorm:"type:varchar(50);unique;not null"`
	Password string `gorm:"type:varchar(50);not null"`
	Pictures string `gorm:"type:text"`
	Events   []eRepo.Event
	Comments []cRepo.Comment
	Payments []pRepo.Payment
}
