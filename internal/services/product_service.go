package services

import (
	"errors"

	"github.com/simpleshop/internal/models"
	"github.com/simpleshop/internal/repository"
)

type ProductInput struct {
	Name        string  `json:"name" binding:"required,min=2"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type ProductService interface {
	Create(input ProductInput) (*models.Product, error)
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Update(id uint, input ProductInput) (*models.Product, error)
	Delete(id uint) error
}

type productService struct {
	repo         repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductService(repo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductService {
	return &productService{repo, categoryRepo}
}

func (s *productService) Create(input ProductInput) (*models.Product, error) {
	_, err := s.categoryRepo.FindByID(input.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	product := &models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		CategoryID:  input.CategoryID,
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return s.repo.FindByID(product.ID)
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Update(id uint, input ProductInput) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}

	_, err = s.categoryRepo.FindByID(input.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	return s.repo.FindByID(product.ID)
}

func (s *productService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}
	return s.repo.Delete(id)
}