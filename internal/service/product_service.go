package service

import (
	"database/sql"
	"errors"
	"simpleshop/internal/model"
	"simpleshop/internal/repository"
)

type ProductService struct {
	productRepo  *repository.ProductRepository
	categoryRepo *repository.CategoryRepository
}

func NewProductService(pr *repository.ProductRepository, cr *repository.CategoryRepository) *ProductService {
	return &ProductService{productRepo: pr, categoryRepo: cr}
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.productRepo.FindAll()
}

func (s *ProductService) GetByID(id int) (*model.Product, error) {
	p, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, errors.New("product not found")
	}
	return p, nil
}

func (s *ProductService) Create(req *model.ProductRequest) (*model.Product, error) {
	cat, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, errors.New("category not found")
	}
	return s.productRepo.Create(req)
}

func (s *ProductService) Update(id int, req *model.ProductRequest) (*model.Product, error) {
	cat, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, errors.New("category not found")
	}
	p, err := s.productRepo.Update(id, req)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, errors.New("product not found")
	}
	return p, nil
}

func (s *ProductService) Delete(id int) error {
	err := s.productRepo.Delete(id)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("product not found")
	}
	return err
}