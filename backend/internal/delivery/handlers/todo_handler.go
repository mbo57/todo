package handlers

import (
	"app/internal/domain"
	"app/internal/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type TodoHandler interface {
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHandlerImpl struct {
	tu usecase.TodoUsecase
}

func NewTodoHandler(tu usecase.TodoUsecase) TodoHandler {
	return &todoHandlerImpl{tu}
}

func (th *todoHandlerImpl) GetAllTodos(c echo.Context) error {
	todos, err := th.tu.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func (th *todoHandlerImpl) GetTodoById(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("todoId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	todo, err := th.tu.GetTodoById(todoId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (th *todoHandlerImpl) CreateTodo(c echo.Context) error {
	todo := domain.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	now := time.Now()
	todo.Created = domain.SetTime(now)
	todo.Updated = domain.SetTime(now)
	if err := th.tu.CreateTodo(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (th *todoHandlerImpl) UpdateTodo(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("todoId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todo := domain.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	beforeTodo, err := th.tu.GetTodoById(todoId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	now := time.Now()
	todo.Created = beforeTodo.Created
	todo.Updated = domain.SetTime(now)
	todo.Id = todoId
	if err := th.tu.UpdateTodo(todoId, todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (th *todoHandlerImpl) DeleteTodo(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("todoId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := th.tu.DeleteTodo(todoId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
