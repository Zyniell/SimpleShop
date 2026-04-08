package handler

import (
	"simpleshop/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine   *gin.Engine
	user     *UserHandler
	category *CategoryHandler
	product  *ProductHandler
	order    *OrderHandler
}

func NewRouter(
	user *UserHandler,
	category *CategoryHandler,
	product *ProductHandler,
	order *OrderHandler,
) *Router {
	return &Router{
		engine:   gin.Default(),
		user:     user,
		category: category,
		product:  product,
		order:    order,
	}
}

func (r *Router) Setup() *gin.Engine {
	api := r.engine.Group("/api")

	// ── User routes (public) ─────────────────────────────────────
	users := api.Group("/users")
	{
		users.POST("/register", r.user.Register)
		users.POST("/login", r.user.Login)
	}

	// ── Category routes ──────────────────────────────────────────
	categories := api.Group("/categories")
	{
		categories.GET("", r.category.GetAll)
		categories.GET("/:id", r.category.GetByID)

		// Protected
		categories.Use(middleware.AuthMiddleware())
		categories.POST("", r.category.Create)
		categories.PUT("/:id", r.category.Update)
		categories.DELETE("/:id", r.category.Delete)
	}

	// ── Product routes ───────────────────────────────────────────
	products := api.Group("/products")
	{
		products.GET("", r.product.GetAll)
		products.GET("/:id", r.product.GetByID)

		// Protected
		products.Use(middleware.AuthMiddleware())
		products.POST("", r.product.Create)
		products.PUT("/:id", r.product.Update)
		products.DELETE("/:id", r.product.Delete)
	}

	// ── Order routes (all protected) ─────────────────────────────
	orders := api.Group("/orders")
	orders.Use(middleware.AuthMiddleware())
	{
		orders.POST("/checkout", r.order.Checkout)
		orders.GET("/my-orders", r.order.GetMyOrders)
	}

	return r.engine
}
