package service

import (
	"errors"
	"simpleshop/internal/model"
	"simpleshop/internal/repository"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	productRepo *repository.ProductRepository
}

func NewOrderService(or *repository.OrderRepository, pr *repository.ProductRepository) *OrderService {
	return &OrderService{orderRepo: or, productRepo: pr}
}

func (s *OrderService) Checkout(userID int, req *model.CheckoutRequest) (*model.Order, error) {
	product, err := s.productRepo.FindByID(req.ProductID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	if product.Stock < req.Quantity {
		return nil, errors.New("insufficient stock")
	}

	if err := s.productRepo.DeductStock(req.ProductID, req.Quantity); err != nil {
		return nil, err
	}

	totalPrice := product.Price * float64(req.Quantity)
	order, err := s.orderRepo.Create(userID, req.ProductID, req.Quantity, totalPrice)
	if err != nil {
		return nil, err
	}

	order.Product = product
	return order, nil
}

func (s *OrderService) GetMyOrders(userID int) ([]model.Order, error) {
	return s.orderRepo.FindByUserID(userID)
}