package repository

import (
	"github.com/MuShaf-NMS/Skyshi-Test/entity"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAll() ([]entity.Todo, error)
	GetAllByActivityGroupID(activityGroupID int) ([]entity.Todo, error)
	Create(todo *entity.Todo) error
	GetOne(id int) (entity.Todo, error)
	Update(todo *entity.Todo, id int) error
	Delete(id int) error
}

type todoRepository struct {
	db *gorm.DB
}

func (tr *todoRepository) GetAll() ([]entity.Todo, error) {
	var todos []entity.Todo
	err := tr.db.Find(&todos).Error
	return todos, err
}

func (tr *todoRepository) GetAllByActivityGroupID(activityGroupID int) ([]entity.Todo, error) {
	var todos []entity.Todo
	err := tr.db.Where(&entity.Todo{ActivityGroupID: activityGroupID}).Find(&todos).Error
	return todos, err
}

func (tr *todoRepository) Create(todo *entity.Todo) error {
	err := tr.db.Create(todo).Error
	return err
}

func (tr *todoRepository) GetOne(id int) (entity.Todo, error) {
	var todo entity.Todo
	err := tr.db.Where(&entity.Todo{ID: id}).First(&todo).Error
	return todo, err
}

func (tr *todoRepository) Update(todo *entity.Todo, id int) error {
	err := tr.db.Where(&entity.Todo{ID: id}).Updates(todo).Error
	return err
}

func (tr *todoRepository) Delete(id int) error {
	err := tr.db.Where(&entity.Todo{ID: id}).Delete(entity.Todo{}).Error
	return err
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}
