package repository

import (
	"app/internal/domain"
	"fmt"

	"gorm.io/gorm"
)

type TodoRepositry interface {
	GetAllTodos(todos *domain.Todos) error
	GetTodoById(todoId int, todo *domain.Todo) error
	CreateTodo(todo *domain.Todo) error
	UpdateTodo(todoId int, todo *domain.Todo) error
	DeleteTodo(todoId int) error
}

type todoRepositryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepositry {
	return &todoRepositryImpl{db}
}

func (tr *todoRepositryImpl) GetAllTodos(todos *domain.Todos) error {
	if err := tr.db.Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepositryImpl) GetTodoById(todoId int, todo *domain.Todo) error {
	if err := tr.db.Where("id = ?", todoId).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepositryImpl) CreateTodo(todo *domain.Todo) error {
	if err := tr.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepositryImpl) UpdateTodo(todoId int, todo *domain.Todo) error {
	if err := tr.db.Model(todo).Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepositryImpl) DeleteTodo(todoId int) error {
	res := tr.db.Model(&domain.Todo{}).Delete("id = ?", todoId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
