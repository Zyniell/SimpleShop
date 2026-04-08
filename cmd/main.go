package main

import (
	"log"

	"github.com/simpleshop/internal/config"
	"github.com/simpleshop/internal/database"
	"github.com/simpleshop/internal/router"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	database.Migrate(db)

	r := router.Setup(db, cfg)

	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}