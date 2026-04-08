package main

import (
	"log"
	"os"
	"simpleshop/config"
	"simpleshop/internal/handler"
	"simpleshop/internal/repository"
	"simpleshop/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env (ignored on Railway — env vars are set via dashboard)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Database
	db := config.InitDB()
	defer db.Close()

	// Repositories
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	// Services
	userSvc := service.NewUserService(userRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	productSvc := service.NewProductService(productRepo, categoryRepo)
	orderSvc := service.NewOrderService(orderRepo, productRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userSvc)
	categoryHandler := handler.NewCategoryHandler(categorySvc)
	productHandler := handler.NewProductHandler(productSvc)
	orderHandler := handler.NewOrderHandler(orderSvc)

	// Router
	router := handler.NewRouter(userHandler, categoryHandler, productHandler, orderHandler)
	engine := router.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("SimpleShop API running on port %s", port)
	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
