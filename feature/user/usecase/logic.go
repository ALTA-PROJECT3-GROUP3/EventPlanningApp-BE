package usecase

import (
	"errors"
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
		}
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
		return user.Core{}, errors.New("terjadi permasalahan server")
	}

	return result, nil
}
