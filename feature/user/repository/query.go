package repository

import (
	"errors"

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

func (um *userModel) Login(username string, password string) (user.Core, error) {
	inputUser := User{}

	if username == "" {
		log.Error("username login is blank")
		return user.Core{}, errors.New("data does not exist")
	}

	if err := um.db.Where("username = ?", username).First(&inputUser).Error; err != nil {
		log.Error("error occurs on select users login", err.Error())
		return user.Core{}, err
	}

	if err := helper.VerifyPassword(inputUser.Password, password); err != nil {
		log.Error("user input for password is wrong", err.Error())
		return user.Core{}, errors.New("wrong password")
	}

	return user.Core{
		ID:       inputUser.Model.ID,
		Name:     inputUser.Name,
		Username: inputUser.Username,
	}, nil
}
