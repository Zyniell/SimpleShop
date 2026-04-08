package services

import (
	"errors"

	"github.com/simpleshop/internal/models"
	"github.com/simpleshop/internal/repository"
)

type CheckoutInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}

type OrderService interface {
	Checkout(userID uint, input CheckoutInput) (*models.Order, error)
	GetUserOrders(userID uint) ([]models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
}

type orderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) OrderService {
	return &orderService{orderRepo, productRepo}
}

func (s *orderService) Checkout(userID uint, input CheckoutInput) (*models.Order, error) {
	product, err := s.productRepo.FindByID(input.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if product.Stock < input.Quantity {
		return nil, errors.New("insufficient stock")
	}

	order := &models.Order{
		UserID:     userID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		TotalPrice: product.Price * float64(input.Quantity),
		Status:     "pending",
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	// Update stock
	product.Stock -= input.Quantity
	s.productRepo.Update(product)

	return s.orderRepo.FindByID(order.ID)
}

func (s *orderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.orderRepo.FindByUserID(userID)
}

func (s *orderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.orderRepo.FindByID(id)
}