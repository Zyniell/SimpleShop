package model

import "time"

type Order struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	ProductID  int       `json:"product_id"`
	Product    *Product  `json:"product,omitempty"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type CheckoutRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,gt=0"`
}