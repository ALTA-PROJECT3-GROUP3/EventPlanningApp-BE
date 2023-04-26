package usecase_test

import (
	"errors"
	"testing"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/mocks"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)
	t.Run("Login Succes", func(t *testing.T) {
		repo.On("Login", "kristain123", "alta123").Return(user.Core{ID: 1, Name: "Kristain Putra", Username: "kristain123"}, nil).Once()
		result, err := ul.LogInLogic("kristain123", "alta123")

		assert.Nil(t, err)
		assert.Equal(t, "kristain123", result.Username)
		assert.Empty(t, result.Password)
		repo.AssertExpectations(t)
	})
	t.Run("Username blank", func(t *testing.T) {
		repo.On("Login", "", "alta123").Return(user.Core{},
			errors.New("data does not exist")).Once()
		result, err := ul.LogInLogic("", "alta123")

		if assert.Error(t, err) {
			assert.Equal(t, "username cannot be blank", err.Error())
		}
		assert.Empty(t, result)
		repo.AssertExpectations(t)
	})
	t.Run("Wrong Password", func(t *testing.T) {
		repo.On("Login", "kristain123", "alta7988").Return(user.Core{},
			errors.New("user input for password is wrong")).Once()
		result, err := ul.LogInLogic("kristain123", "alta7988")

		if assert.Error(t, err) {
			assert.Equal(t, "password is wrong", err.Error())
		}
		assert.Empty(t, result)
		repo.AssertExpectations(t)
	})
	t.Run("Internal DB Error", func(t *testing.T) {
		repo.On("Login", "kristaink1237988", "alta7988").Return(user.Core{}, errors.New("too much")).Once()
		result, err := ul.LogInLogic("kristaink1237988", "alta7988")

		if assert.Error(t, err) {
			assert.Equal(t, "internal server error", err.Error())
		}
		assert.Empty(t, result)
		repo.AssertExpectations(t)
	})
}

func TestRegisterUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)
	var newUser user.Core
	newUser.Name = "Kristain Putra"
	newUser.Password = "alta1237988<p>"
	newUser.Username = "kristain0998"

	t.Run("server DB Error", func(t *testing.T) {
		repo.On("InsertUser", newUser).Return(errors.New("column")).Once()
		err := ul.RegisterUser(newUser)

		if assert.Error(t, err) {
			assert.Equal(t, "server error", err.Error())
		}
		repo.AssertExpectations(t)
	})
	t.Run("Invalid Value", func(t *testing.T) {
		repo.On("InsertUser", newUser).Return(errors.New("value")).Once()
		err := ul.RegisterUser(newUser)

		if assert.Error(t, err) {
			assert.Equal(t, "invalid value", err.Error())
		}
		repo.AssertExpectations(t)
	})
	t.Run("Server error", func(t *testing.T) {
		repo.On("InsertUser", newUser).Return(errors.New("too much")).Once()
		err := ul.RegisterUser(newUser)

		if assert.Error(t, err) {
			assert.Equal(t, "server error", err.Error())
		}
		repo.AssertExpectations(t)
	})
	t.Run("Password too short", func(t *testing.T) {
		repo.On("InsertUser", newUser).Return(errors.New("too short")).Once()
		err := ul.RegisterUser(newUser)

		if assert.Error(t, err) {
			assert.Equal(t, "invalid password length", err.Error())
		}
		repo.AssertExpectations(t)
	})
	t.Run("Succes Register", func(t *testing.T) {
		repo.On("InsertUser", newUser).Return(nil).Once()
		err := ul.RegisterUser(newUser)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUserProfile(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)
	existUser := user.Core{
		ID:       1,
		Name:     "Kristain Putra",
		Email:    "kristainputra98@gmail.com",
		Picture:  "https://res.cloudinary.com/dc0wgpho2/image/upload/v1682413943/Home/EventPlanningApp/users/etzwu5atihgqowt93nn0.jpg",
		Username: "kristain09",
	}
	t.Run("Succes to user profile", func(t *testing.T) {
		repo.On("GetUserById", existUser.ID).Return(user.Core{
			Name:    existUser.Name,
			Email:   existUser.Email,
			Picture: existUser.Picture,
		}, nil).Once()
		res, err := ul.UserProfileLogic(existUser.ID)
		assert.Nil(t, err)
		if assert.NotEmpty(t, res) {
			assert.Equal(t, existUser.Name, res.Name)
			assert.Equal(t, existUser.Email, res.Email)
			assert.Equal(t, existUser.Picture, res.Picture)
		}
		repo.AssertExpectations(t)
	})
	t.Run("failed to find user", func(t *testing.T) {
		repo.On("GetUserById", uint(2)).Return(user.Core{}, errors.New("too much"))
		res, err := ul.UserProfileLogic(uint(2))
		if assert.NotNil(t, err) {
			assert.Equal(t, "internal server error", err.Error())
		}
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUserLogic(t *testing.T) {
	// create a new mock repository
	repo := new(mocks.Repository)
	ul := usecase.New(repo)

	repo.On("DeleteUser", uint(1)).Return(errors.New("finding user"))
	err := ul.DeleteUserLogic(1)
	assert.EqualError(t, err, "bad request, user not found")

	repo.On("DeleteUser", uint(2)).Return(errors.New("cannot delete user"))
	err = ul.DeleteUserLogic(2)
	assert.EqualError(t, err, "internal server error, cannot delete user")

	repo.On("DeleteUser", uint(3)).Return(errors.New("unknown error"))
	err = ul.DeleteUserLogic(3)
	assert.EqualError(t, err, "unknown error")

	repo.AssertExpectations(t)
}

func TestUpdateProfile(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockUseCase := usecase.New(mockRepo)

	mockRepo.On("UpdateProfile", uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", mock.AnythingOfType("*multipart.FileHeader")).Return(nil).Once()
	err := mockUseCase.UpdateProfileLogic(uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", nil)
	assert.NoError(t, err)

	mockRepo.On("UpdateProfile", uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", mock.AnythingOfType("*multipart.FileHeader")).Return(errors.New("open /path/to/file: permission denied")).Once()
	err = mockUseCase.UpdateProfileLogic(uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", nil)
	assert.EqualError(t, err, "user photo are not allowed")

	mockRepo.On("UpdateProfile", uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", mock.AnythingOfType("*multipart.FileHeader")).Return(errors.New("upload file in path")).Once()
	err = mockUseCase.UpdateProfileLogic(uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", nil)
	assert.EqualError(t, err, "cannot upload file in path")

	mockRepo.On("UpdateProfile", uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", mock.AnythingOfType("*multipart.FileHeader")).Return(errors.New("hashing password")).Once()
	err = mockUseCase.UpdateProfileLogic(uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", nil)
	assert.EqualError(t, err, "is invalid")

	mockRepo.On("UpdateProfile", uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", mock.AnythingOfType("*multipart.FileHeader")).Return(errors.New("no rows affected")).Once()
	err = mockUseCase.UpdateProfileLogic(uint(1), "Kristain Putra", "kristainputra98@gmail.com", "alta123", nil)
	assert.EqualError(t, err, "data is up to date")

	mockRepo.AssertExpectations(t)
}
