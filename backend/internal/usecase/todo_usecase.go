package usecase

import (
	"app/internal/domain"
	"app/internal/repository"
)

type TodoUsecase interface {
	GetAllTodos() (domain.Todos, error)
	GetTodoById(todoId int) (domain.Todo, error)
	CreateTodo(todo domain.Todo) error
	UpdateTodo(todoId int, todo domain.Todo) error
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

func (tu *todoUsecaseImpl) GetTodoById(todoId int) (domain.Todo, error) {
	todo := domain.Todo{}
	if err := tu.tr.GetTodoById(todoId, &todo); err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (tu *todoUsecaseImpl) CreateTodo(todo domain.Todo) error {
	if err := tu.tr.CreateTodo(&todo); err != nil {
		return err
	}
	return nil
}

func (tu *todoUsecaseImpl) UpdateTodo(todoId int, todo domain.Todo) error {
	if err := tu.tr.UpdateTodo(todoId, &todo); err != nil {
		return err
	}
	return nil
}
