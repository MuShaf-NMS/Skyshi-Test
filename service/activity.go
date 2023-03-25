package service

import (
	"fmt"

	"github.com/MuShaf-NMS/Skyshi-Test/dto"
	"github.com/MuShaf-NMS/Skyshi-Test/entity"
	"github.com/MuShaf-NMS/Skyshi-Test/helper"
	"github.com/MuShaf-NMS/Skyshi-Test/repository"
)

type ActivityService interface {
	GetAll() ([]entity.Activity, error)
	Create(createActivity dto.CreateActivity) (entity.Activity, error)
	GetOne(id int) (entity.Activity, error)
	Update(updateActivity dto.UpdateActivity, id int) (entity.Activity, error)
	Delete(id int) error
}

type activityService struct {
	repository repository.ActivityRepository
}

func (as *activityService) GetAll() ([]entity.Activity, error) {
	activities, err := as.repository.GetAll()
	return activities, err
}

func (as *activityService) Create(createActivity dto.CreateActivity) (entity.Activity, error) {
	activity := entity.Activity{
		Title: createActivity.Title,
		Email: createActivity.Email,
	}
	err := as.repository.Create(&activity)
	if err != nil {
		return activity, helper.NewError(422, "Failed", "Failed to create new activity")
	}
	return activity, nil
}

func (as *activityService) GetOne(id int) (entity.Activity, error) {
	activity, err := as.repository.GetOne(id)
	if err != nil {
		return activity, helper.NewError(404, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
	}
	return activity, nil
}

func (as *activityService) Update(updateActivity dto.UpdateActivity, id int) (entity.Activity, error) {
	activity := entity.Activity{
		Title: updateActivity.Title,
	}
	err := as.repository.Update(&activity, id)
	if err != nil {
		return activity, helper.NewError(422, "Failed", "Failed to create new activity")
	}
	activity, err = as.GetOne(id)
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (as *activityService) Delete(id int) error {
	_, err := as.GetOne(id)
	if err != nil {
		return err
	}
	err = as.repository.Delete(id)
	if err != nil {
		return helper.NewError(404, "Not Found", fmt.Sprintf("Activity with ID %d Not Found", id))
	}
	return nil
}

func NewActivityService(repository repository.ActivityRepository) ActivityService {
	return &activityService{
		repository: repository,
	}
}
