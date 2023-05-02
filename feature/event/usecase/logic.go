package usecase

import (
	"mime/multipart"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/feature/event"
	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/utils/helper"
	"github.com/go-playground/validator/v10"
)

type eventLogic struct {
	data event.Repository
	vld  *validator.Validate
}

func New(repo event.Repository) event.UseCase {
	return &eventLogic{
		data: repo,
		vld:  validator.New(),
	}
}

// Update implements event.UseCase
func (uc *eventLogic) Update(userId uint, id uint, updateEvent event.Core, file *multipart.FileHeader) error {
	if file != nil {
		file, _ := file.Open()
		uploadURL, err := helper.UploadFile(file, "/events")
		if err != nil {
			return err
		}
		updateEvent.Pictures = uploadURL[0]
	}

	errUpdate := uc.data.Update(userId, id, updateEvent)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

// GetEventById implements event.UseCase
func (uc *eventLogic) GetEventById(id uint) (event.Core, error) {
	data, err := uc.data.GetEventById(id)
	return data, err
}

// MyEvent implements event.UseCase
func (uc *eventLogic) MyEvent(userId uint, page int) ([]event.Core, error) {
	limit := 10
	offset := (page - 1) * limit
	data, err := uc.data.MyEvent(userId, limit, offset)
	return data, err
}

// GetAll implements event.UseCase
func (uc *eventLogic) GetAll(page int, name string) ([]event.Core, error) {
	limit := 10
	offset := (page - 1) * limit
	data, err := uc.data.SelectAll(limit, offset, name)
	return data, err
}

// Add implements event.UseCase
func (uc *eventLogic) Add(newEvent event.Core, file *multipart.FileHeader) error {
	errValidate := uc.vld.Struct(newEvent)
	if errValidate != nil {
		return errValidate
	}

	if file != nil {
		file, _ := file.Open()
		uploadURL, err := helper.UploadFile(file, "/events")
		if err != nil {
			return err
		}
		newEvent.Pictures = uploadURL[0]
	}

	errInsert := uc.data.Insert(newEvent)
	if errInsert != nil {
		return errInsert
	}
	return nil
}
