package handlers

import (
	"app/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TodoHandler interface {
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	// CreateTodo(c echo.Context) error
	// UpdateTodo(c echo.Context) error
	// DeleteTodo(c echo.Context) error
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
