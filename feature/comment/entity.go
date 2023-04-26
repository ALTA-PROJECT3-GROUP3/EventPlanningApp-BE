package comment

import "github.com/labstack/echo/v4"

type Core struct {
	UserID  uint
	EventID uint
	Comment string
}

type Handler interface {
	CreateCommentHandler() echo.HandlerFunc
}

type UseCase interface {
	CreateCommentLogic(Core) error
}

type Repository interface {
	InsertComment(Core) error
}
