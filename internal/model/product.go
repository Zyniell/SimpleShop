package model

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  int       `json:"category_id"`
	Category    *Category `json:"category,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
	CategoryID  int     `json:"category_id" binding:"required"`
}