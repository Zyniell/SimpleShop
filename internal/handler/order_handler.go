package handler

import (
	"net/http"
	"simpleshop/internal/model"
	"simpleshop/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(s *service.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req model.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.Checkout(userID.(int), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Checkout successful",
		"data":    order,
	})
}

func (h *OrderHandler) GetMyOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	orders, err := h.service.GetMyOrders(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	if orders == nil {
		orders = []model.Order{}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Orders fetched successfully",
		"data":    orders,
	})
}
