package service

import (
	"fmt"

	"github.com/MuShaf-NMS/Skyshi-Test/dto"
	"github.com/MuShaf-NMS/Skyshi-Test/entity"
	"github.com/MuShaf-NMS/Skyshi-Test/helper"
	"github.com/MuShaf-NMS/Skyshi-Test/repository"
)

type TodoService interface {
	GetAll(activity_group_id int) ([]entity.Todo, error)
	Create(createTodo dto.CreateTodo) (entity.Todo, error)
	GetOne(id int) (entity.Todo, error)
	Update(upateTodo dto.UpdateTodo, id int) (entity.Todo, error)
	Delete(id int) error
}

type todoService struct {
	repository repository.TodoRepository
}

func (ts *todoService) GetAll(activity_group_id int) ([]entity.Todo, error) {
	if activity_group_id == 0 {
		todos, err := ts.repository.GetAll()
		return todos, err
	}
	todos, err := ts.repository.GetAllByActivityGroupID(activity_group_id)
	return todos, err
}

func (ts *todoService) Create(createTodo dto.CreateTodo) (entity.Todo, error) {
	isActive := true
	if createTodo.IsActive != nil {
		isActive = *createTodo.IsActive
	}
	todo := entity.Todo{
		Title:           createTodo.Title,
		ActivityGroupID: createTodo.ActivityGroupID,
		IsActive:        isActive,
	}
	err := ts.repository.Create(&todo)
	if err != nil {
		return todo, helper.NewError(422, "Failed", "Failed to create new todo")
	}
	return todo, nil
}

func (ts *todoService) GetOne(id int) (entity.Todo, error) {
	todo, err := ts.repository.GetOne(id)
	if err != nil {
		return todo, helper.NewError(404, "Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
	}
	return todo, nil
}

func (ts *todoService) Update(updateTodo dto.UpdateTodo, id int) (entity.Todo, error) {
	todo := entity.Todo{
		Title:    updateTodo.Title,
		Priority: updateTodo.Priority,
		IsActive: updateTodo.IsActive,
		Status:   updateTodo.Status,
	}
	err := ts.repository.Update(&todo, id)
	if err != nil {
		return todo, helper.NewError(422, "Failed", "Failed to create new todo")
	}
	todo, err = ts.GetOne(id)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (ts *todoService) Delete(id int) error {
	_, err := ts.GetOne(id)
	if err != nil {
		return err
	}
	err = ts.repository.Delete(id)
	if err != nil {
		return helper.NewError(422, "Failed", "Failed to create new todo")
	}
	return nil
}

func NewTodoService(repository repository.TodoRepository) TodoService {
	return &todoService{
		repository: repository,
	}
}
