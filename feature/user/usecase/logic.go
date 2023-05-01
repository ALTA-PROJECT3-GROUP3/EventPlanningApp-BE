package usecase

import (
	"errors"
	"mime/multipart"
	"strings"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user"
	"github.com/labstack/gommon/log"
)

type userLogic struct {
	u user.Repository
}

func New(r user.Repository) user.UseCase {
	return &userLogic{
		u: r,
	}
}

func (ul userLogic) RegisterUser(newUser user.Core) error {
	if err := ul.u.InsertUser(newUser); err != nil {
		log.Error("error on calling register insert user query", err.Error())
		if strings.Contains(err.Error(), "column") {
			return errors.New("server error")
		} else if strings.Contains(err.Error(), "value") {
			return errors.New("invalid value")
		} else if strings.Contains(err.Error(), "too short") {
			return errors.New("invalid password length")
		}
		return errors.New("server error")
	}
	return nil
}

func (ul userLogic) LogInLogic(username string, password string) (user.Core, error) {
	res, err := ul.u.Login(username, password)
	if err != nil {

		if strings.Contains(err.Error(), "not exist") {
			return user.Core{}, errors.New("username cannot be blank")

		} else if strings.Contains(err.Error(), "wrong") {
			return user.Core{}, errors.New("password is wrong")

		}
		log.Error("error on loginlogic, internal server error", err.Error())
		return user.Core{}, errors.New("internal server error")

	}

	return res, nil
}

func (ul userLogic) UserProfileLogic(id uint) (user.Core, error) {
	result, err := ul.u.GetUserById(id)
	if err != nil {
		log.Error("failed to find user", err.Error())
		return user.Core{}, errors.New("internal server error")
	}

	return result, nil
}

func (ul userLogic) DeleteUserLogic(id uint) error {
	err := ul.u.DeleteUser(id)
	if err != nil {
		log.Error("failed on calling deleteuser query")
		if strings.Contains(err.Error(), "finding user") {
			log.Error("error on finding user (not found)")
			return errors.New("bad request, user not found")
		} else if strings.Contains(err.Error(), "cannot delete") {
			log.Error("error on delete user")
			return errors.New("internal server error, cannot delete user")
		}
		log.Error("error in delete user (else)")
		return err
	}
	return nil
}

func (ul *userLogic) UpdateProfileLogic(id uint, name string, email string, password string, picture *multipart.FileHeader) error {
	if err := ul.u.UpdateProfile(id, name, email, password, picture); err != nil {
		log.Error("failed on calling updateprofile query")
		if strings.Contains(err.Error(), "open") {
			log.Error("errors occurs on opening picture file")
			return errors.New("user photo are not allowed")
		} else if strings.Contains(err.Error(), "upload file in path") {
			log.Error("upload file in path are error")
			return errors.New("cannot upload file in path")
		} else if strings.Contains(err.Error(), "hashing password") {
			log.Error("hashing password error")
			return errors.New("is invalid")
		} else if strings.Contains(err.Error(), "affected") {
			log.Error("no rows affected on update user")
			return errors.New("data is up to date")
		}
		return err
	}
	return nil
}
