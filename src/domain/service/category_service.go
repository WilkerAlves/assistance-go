package service

import (
	"errors"
	"fmt"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/interface/repository"
)

const (
	MessageValidationCategoryName = "the category name already exists"
)

type ICategoryService interface {
	Create(category entity.Category) error
	Update(category entity.Category) error
	GetById(id string) (*entity.Category, error)
	GetByName(name string) (*entity.Category, error)
	GetAll() ([]*entity.Category, error)
}

type categoryService struct {
	repo repository.ICategoryRepository
}

func (s *categoryService) Create(category entity.Category) error {
	if cat, _ := s.repo.FindByName(category.GetName()); cat != nil {
		return errors.New(MessageValidationCategoryName)
	}

	err := s.repo.Create(category)
	if err != nil {
		return fmt.Errorf("error while create category, %w", err)
	}

	return nil
}

func (s *categoryService) Update(category entity.Category) error {
	if len(category.GetName()) < 1 {
		return errors.New("the category ai is empty")
	}

	cat, err := s.repo.FindByName(category.GetName())
	if err != nil {
		return err
	}

	if cat != nil && cat.GetID() != category.GetID() {
		return errors.New(MessageValidationCategoryName)
	}

	return s.repo.Update(category)
}

func (s *categoryService) GetById(id string) (*entity.Category, error) {
	if len(id) < 1 {
		return nil, errors.New("id is empty")
	}
	category, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetByName(name string) (*entity.Category, error) {
	if len(name) < 1 {
		return nil, errors.New("name is empty")
	}
	category, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetAll() ([]*entity.Category, error) {
	category, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func NewCategoryService(repo repository.ICategoryRepository) *categoryService {
	return &categoryService{repo: repo}
}
