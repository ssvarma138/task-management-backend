package main

import (
	"log"
	"task-management-backend/internal/api"
	"task-management-backend/internal/config"
	"task-management-backend/internal/storage"
)

func main() {
	cfg := config.Load()
	db, err := storage.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := api.SetupRouter(db)
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}