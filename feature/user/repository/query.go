package repository

import (
	"errors"
	"mime/multipart"

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

func (um *userModel) GetUserById(id uint) (user.Core, error) {
	var res user.Core
	if err := um.db.Table("users").Select("name, email, pictures").Where("id = ?", id).First(&res).Error; err != nil {
		log.Error("error occurs in finding user profile", err.Error())
		return user.Core{}, err
	}

	return res, nil
}

func (um *userModel) DeleteUser(id uint) error {
	userToDelete := &User{}
	if err := um.db.First(userToDelete, id).Error; err != nil {
		log.Error("Error in finding user id")
		return errors.New("error in finding user")
	}

	if err := um.db.Delete(userToDelete).Error; err != nil {
		log.Error("cannot deleete user")
		return err
	}

	return nil
}

func (um *userModel) UpdateProfile(id uint, name string, email string, password string, picture *multipart.FileHeader) error {
	var UpdateUser User
	if picture != nil {
		file, err := picture.Open()
		if err != nil {
			log.Errorf("error occurs on open picture %v", err)
			return errors.New("error on open picture")
		}

		uploadURL, err := helper.UploadFile(file, "/users")
		if err != nil {
			log.Errorf("error occurs on uploadFile in path %v", err)
			return errors.New("error on upload file in path")
		}
		UpdateUser.Picture = uploadURL[0]
	}

	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		log.Error("error occurs on hashing password", err.Error())
		return errors.New("hashing password failed")
	}

	UpdateUser.ID = id
	UpdateUser.Name = name
	UpdateUser.Email = email
	UpdateUser.Password = hashedPassword

	tx := um.db.Where("id = ?", id).First(&UpdateUser)
	if tx.RowsAffected < 1 {
		log.Error("there is no column to change on update user")
		return errors.New("no data affected")
	}
	if tx.Error != nil {
		log.Error("error on update user")
		return err
	}

	return nil
}
