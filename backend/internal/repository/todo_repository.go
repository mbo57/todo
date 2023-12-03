package repository

import (
	"app/internal/domain"

	"gorm.io/gorm"
)

type TodoRepositry interface {
	GetAllTodos(todos *domain.Todos) error
	// GetTodoById(todoId int) error
	// CreateTodo(todo *domain.Todo) error
	// UpdateTodo(todo *domain.Todo) error
	// DeleteTodo(todoId int) error
}

type todoRepositryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepositry {
	return &todoRepositryImpl{db}
}

func (tr *todoRepositryImpl) GetAllTodos(todos *domain.Todos) error {
	if err := tr.db.Find(&todos).Error; err != nil {
		return err
	}
	return nil
}
