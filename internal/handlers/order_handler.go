package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/simpleshop/internal/services"
	"github.com/simpleshop/pkg/utils"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

// POST /api/orders/checkout  [PROTECTED]
func (h *OrderHandler) Checkout(c *gin.Context) {
	userID, _ := c.Get("userID")

	var input services.CheckoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	order, err := h.service.Checkout(userID.(uint), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Order placed successfully", order))
}

// GET /api/orders  [PROTECTED] - get current user's orders
func (h *OrderHandler) GetMyOrders(c *gin.Context) {
	userID, _ := c.Get("userID")

	orders, err := h.service.GetUserOrders(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Orders fetched", orders))
}

// GET /api/orders/:id  [PROTECTED]
func (h *OrderHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
		return
	}

	order, err := h.service.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Order not found"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Order fetched", order))
}