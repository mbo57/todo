package routes

import (
	"app/internal/delivery/handlers"

	"github.com/labstack/echo"
)

func NewTodoRouter(th handlers.TodoHandler) *echo.Echo {
	e := echo.New()
	t := e.Group("todos")
	t.GET("", th.GetAllTodos)

	return e
}
