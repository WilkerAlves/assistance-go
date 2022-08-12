package service

import (
	"errors"
	"fmt"

	"github.com/WilkerAlves/assistance-go/src/domain/dto"
	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/interfaces"
)

const (
	MessageValidationCategoryName = "the category name already exists"
)

type categoryService struct {
	repo interfaces.ICategoryRepository
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

func (s *categoryService) GetAll(filters *dto.CategoryFiltersDTO) ([]entity.Category, error) {
	var categories []entity.Category
	var err error

	if filters == nil {
		categories, err = s.repo.FindAll(nil)
		if err != nil {
			return categories, err
		}
		return categories, nil
	}

	categories, err = s.repo.FindAll(filters.Active)
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func NewCategoryService(repo interfaces.ICategoryRepository) *categoryService {
	return &categoryService{repo: repo}
}
