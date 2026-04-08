package service

import (
	"database/sql"
	"errors"
	"simpleshop/internal/model"
	"simpleshop/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) GetByID(id int) (*model.Category, error) {
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, errors.New("category not found")
	}
	return cat, nil
}

func (s *CategoryService) Create(req *model.CategoryRequest) (*model.Category, error) {
	return s.repo.Create(req.Name)
}

func (s *CategoryService) Update(id int, req *model.CategoryRequest) (*model.Category, error) {
	cat, err := s.repo.Update(id, req.Name)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, errors.New("category not found")
	}
	return cat, nil
}

func (s *CategoryService) Delete(id int) error {
	err := s.repo.Delete(id)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("category not found")
	}
	return err
}