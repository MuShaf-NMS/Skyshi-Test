package repository

import (
	"github.com/MuShaf-NMS/Skyshi-Test/entity"
	"gorm.io/gorm"
)

type ActivityRepository interface {
	GetAll() ([]entity.Activity, error)
	Create(activity *entity.Activity) error
	GetOne(id int) (entity.Activity, error)
	Update(activity *entity.Activity, id int) error
	Delete(id int) error
}

type activityRepository struct {
	db *gorm.DB
}

func (ar *activityRepository) GetAll() ([]entity.Activity, error) {
	var activities []entity.Activity
	err := ar.db.Find(&activities).Error
	return activities, err
}

func (ar *activityRepository) Create(activity *entity.Activity) error {
	err := ar.db.Create(activity).Error
	return err
}

func (ar *activityRepository) GetOne(id int) (entity.Activity, error) {
	var activity entity.Activity
	err := ar.db.Where(&entity.Activity{ID: id}).First(&activity).Error
	return activity, err
}

func (ar *activityRepository) Update(activity *entity.Activity, id int) error {
	err := ar.db.Where(&entity.Activity{ID: id}).Updates(activity).Error
	return err
}

func (ar *activityRepository) Delete(id int) error {
	err := ar.db.Where(&entity.Activity{ID: id}).Delete(entity.Activity{}).Error
	return err
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{
		db: db,
	}
}
