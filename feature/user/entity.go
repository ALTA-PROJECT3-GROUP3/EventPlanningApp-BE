package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Username string
	Picture  string
}

type Handler interface {
	RegisterHandler() echo.HandlerFunc
	LoginHandler() echo.HandlerFunc
	UserProfileHandler() echo.HandlerFunc
	DeleteUserHandler() echo.HandlerFunc
	UpdateProfileHandler() echo.HandlerFunc
}

type UseCase interface {
	RegisterUser(newUser Core) error
	LogInLogic(username string, password string) (Core, error)
	UserProfileLogic(id uint) (Core, error)
	DeleteUserLogic(id uint) error
	UpdateProfileLogic(id uint, name string, email string, password string, picture *multipart.FileHeader) error
}

type Repository interface {
	InsertUser(newUser Core) error
	Login(username string, password string) (Core, error)
	GetUserById(id uint) (Core, error)
	DeleteUser(id uint) error
	UpdateProfile(id uint, name string, email string, password string, picture *multipart.FileHeader) error
}
