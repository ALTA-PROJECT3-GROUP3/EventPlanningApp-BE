package usecase

import (
	"errors"
	"strings"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/comment"
	"github.com/labstack/gommon/log"
)

type commentLogic struct {
	cm comment.Repository
}

func New(cm comment.Repository) comment.UseCase {
	return &commentLogic{
		cm: cm,
	}
}

func (cl *commentLogic) CreateCommentLogic(newComment comment.Core) error {
	err := cl.cm.InsertComment(newComment)
	if err != nil {

		if strings.Contains(err.Error(), "connect") {
			log.Error("cannot connect to database")
			return errors.New("failed to connect to database")
		}

		if strings.Contains(err.Error(), "table 'events' not found") {
			log.Error("table events not found")
			return errors.New("table 'events' not found")
		}

		if strings.Contains(err.Error(), "event not found") {
			log.Error("event id is does not exist")
			return errors.New("event not found (bad request)")
		}

		if strings.Contains(err.Error(), "table 'comments' not found") {
			log.Error("table 'comments' not found")
			return errors.New("table 'comments' not found")
		}

		log.Error("undelared error in insert comment")
		return errors.New("server error")
	}

	return nil
}
