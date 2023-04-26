package usecase_test

import (
	"testing"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/mocks"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/user/usecase"
)

func TestLogin(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)

	t.Run("Sukses Login", func(t *testing.T){
		repo.On("")

	})
}
