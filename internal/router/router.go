package router

import (
	"github.com/gin-gonic/gin"
	"github.com/simpleshop/internal/config"
	"github.com/simpleshop/internal/handlers"
	"github.com/simpleshop/internal/middleware"
	"github.com/simpleshop/internal/repository"
	"github.com/simpleshop/internal/services"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Repositories
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	// Services
	userSvc := services.NewUserService(userRepo, cfg)
	categorySvc := services.NewCategoryService(categoryRepo)
	productSvc := services.NewProductService(productRepo, categoryRepo)
	orderSvc := services.NewOrderService(orderRepo, productRepo)

	// Handlers
	userHandler := handlers.NewUserHandler(userSvc)
	categoryHandler := handlers.NewCategoryHandler(categorySvc)
	productHandler := handlers.NewProductHandler(productSvc)
	orderHandler := handlers.NewOrderHandler(orderSvc)

	// Auth middleware
	authMiddleware := middleware.AuthMiddleware(cfg)

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "SimpleShop API is running", "version": "1.0.0"})
	})

	api := r.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)
		}

		// Category routes
		categories := api.Group("/categories")
		{
			categories.GET("", categoryHandler.GetAll)
			categories.GET("/:id", categoryHandler.GetByID)
			categories.POST("", authMiddleware, categoryHandler.Create)
			categories.PUT("/:id", authMiddleware, categoryHandler.Update)
			categories.DELETE("/:id", authMiddleware, categoryHandler.Delete)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetAll)
			products.GET("/:id", productHandler.GetByID)
			products.POST("", authMiddleware, productHandler.Create)
			products.PUT("/:id", authMiddleware, productHandler.Update)
			products.DELETE("/:id", authMiddleware, productHandler.Delete)
		}

		// Order routes (all protected)
		orders := api.Group("/orders", authMiddleware)
		{
			orders.POST("/checkout", orderHandler.Checkout)
			orders.GET("", orderHandler.GetMyOrders)
			orders.GET("/:id", orderHandler.GetByID)
		}
	}

	return r
}