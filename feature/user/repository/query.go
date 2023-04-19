package repository

import (
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &userModel{
		db: db,
	}
}

func (um *userModel) InsertUser(newUser user.Core) error {
	inputUser := User{}
	hashedPassword, err := helper.HashPassword(newUser.Password)
	if err != nil {
		log.Error("error occurs on hashing password", err.Error())
		return err
	}

	inputUser.Name = newUser.Name
	inputUser.Email = newUser.Email
	inputUser.Username = newUser.Username
	inputUser.Picture = newUser.Picture
	inputUser.Password = hashedPassword

	if err := um.db.Table("users").Create(&inputUser).Error; err != nil {
		log.Error("error on create table users", err.Error())
		return err
	}

	return nil
}
