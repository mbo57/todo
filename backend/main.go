package main

import (
	"app/internal/delivery/handlers"
	"app/internal/delivery/routes"
	"app/internal/repository"
	"app/internal/repository/database"
	"app/internal/usecase"
	"log"
	"time"
)

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}
	time.Local = jst
	db := database.NewDB()
	todoRepositry := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepositry)
	todoHandler := handlers.NewTodoHandler(todoUsecase)
	e := routes.NewTodoRouter(todoHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
