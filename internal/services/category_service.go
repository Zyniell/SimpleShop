package services

import (
	"errors"

	"github.com/simpleshop/internal/models"
	"github.com/simpleshop/internal/repository"
)

type CategoryInput struct {
	Name string `json:"name" binding:"required,min=2"`
}

type CategoryService interface {
	Create(input CategoryInput) (*models.Category, error)
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	Update(id uint, input CategoryInput) (*models.Category, error)
	Delete(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) Create(input CategoryInput) (*models.Category, error) {
	category := &models.Category{Name: input.Name}
	if err := s.repo.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetByID(id uint) (*models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) Update(id uint, input CategoryInput) (*models.Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}
	category.Name = input.Name
	if err := s.repo.Update(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}
	return s.repo.Delete(id)
}