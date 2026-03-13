package main

import (
	"log"
	"test-go/internal/config"
	"test-go/internal/database"
	"test-go/internal/delivery/http"
	"test-go/internal/repository"
	"test-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.MigrateDB(db); err != nil {
		log.Fatal(err)
	}

	todoRepo := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")

	http.NewTodoHandler(v1, todoUsecase)

	r.Run(":8080")
}