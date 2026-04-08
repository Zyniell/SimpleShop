package handler

import (
	"net/http"
	"simpleshop/internal/model"
	"simpleshop/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	if categories == nil {
		categories = []model.Category{}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categories fetched successfully",
		"data":    categories,
	})
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category fetched successfully",
		"data":    category,
	})
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req model.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
		"data":    category,
	})
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var req model.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.service.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category updated successfully",
		"data":    category,
	})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
