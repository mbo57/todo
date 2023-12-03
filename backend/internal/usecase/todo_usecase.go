package usecase

import (
	"app/internal/domain"
	"app/internal/repository"
)

type TodoUsecase interface {
	GetAllTodos() (domain.Todos, error)
	// GetTodoById(todoId int) (domain.Todo, error)
	// CreateTodo(todo domain.Todo) error
	// UpdateTodo(todo domain.Todo) error
	// DeleteTodo(todoId int) error
}

type todoUsecaseImpl struct {
	tr repository.TodoRepositry
}

func NewTodoUsecase(tr repository.TodoRepositry) TodoUsecase {
	return &todoUsecaseImpl{tr}
}

func (tu *todoUsecaseImpl) GetAllTodos() (domain.Todos, error) {
	todos := domain.Todos{}
	if err := tu.tr.GetAllTodos(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}