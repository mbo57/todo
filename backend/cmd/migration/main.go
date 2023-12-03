package main

import (
	"app/internal/domain"
	"app/internal/repository/database"
	"fmt"
)

func main() {
	dbConn := database.NewDB()
	defer fmt.Println("maigration db")
	defer database.CloseDB(dbConn)
	dbConn.AutoMigrate(&domain.Todo{})
}
