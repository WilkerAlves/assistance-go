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
}

type CategoryService struct {
	Repo repository.ICategoryRepository
}

func (s *CategoryService) Create(category entity.Category) error {
	if cat, _ := s.Repo.FindByName(category.GetName()); cat != nil {
		return errors.New(MessageValidationCategoryName)
	}

	err := s.Repo.Create(category)
	if err != nil {
		return fmt.Errorf("error while create category, %w", err)
	}

	return nil
}

func (s *CategoryService) Update(category entity.Category) error {
	if len(category.GetName()) < 1 {
		return errors.New("the category ai is empty")
	}

	if cat, _ := s.Repo.FindByName(category.GetName()); cat != nil && cat.GetID() != category.GetID() {
		return errors.New(MessageValidationCategoryName)
	}
	return s.Repo.Update(category)
}

func (s *CategoryService) GetById(id string) (*entity.Category, error) {
	if len(id) < 1 {
		return nil, errors.New("id is empty")
	}
	category, err := s.Repo.Find(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) GetByName(name string) (*entity.Category, error) {
	if len(name) < 1 {
		return nil, errors.New("name is empty")
	}
	category, err := s.Repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func NewCategoryService(repo repository.ICategoryRepository) *CategoryService {
	return &CategoryService{Repo: repo}
}
