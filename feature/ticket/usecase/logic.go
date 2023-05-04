package usecase

import (
	"errors"
	"strings"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/ticket"
	"github.com/labstack/gommon/log"
)

type ticketLogic struct {
	tc ticket.Repository
}

func New(tc ticket.Repository) ticket.UseCase {
	return &ticketLogic{
		tc: tc,
	}
}

// Update implements ticket.UseCase
func (tk *ticketLogic) Update(userId uint, id uint, updateTicket ticket.Core) error {
	errUpdate := tk.tc.Update(userId, id, updateTicket)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

// Create implements ticket.UseCase
func (tk *ticketLogic) Create(newTicket ticket.Core) error {
	err := tk.tc.Insert(newTicket)
	if err != nil {

		if strings.Contains(err.Error(), "connect") {
			log.Error("cannot connect to database")
			return errors.New("failed to connect to database")
		}

		if strings.Contains(err.Error(), "table 'tickets' not found") {
			log.Error("table tickets not found")
			return errors.New("table 'tickets' not found")
		}

		if strings.Contains(err.Error(), "ticket not found") {
			log.Error("ticket id is does not exist")
			return errors.New("ticket not found (bad request)")
		}

		if strings.Contains(err.Error(), "table 'tickets' not found") {
			log.Error("table 'tickets' not found")
			return errors.New("table 'tickets' not found")
		}

		log.Error("undelared error in insert ticket")
		return errors.New("server error")
	}

	return nil
}
