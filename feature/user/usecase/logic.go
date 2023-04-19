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
