package mocks

import (
	"errors"
	"fmt"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/interface/repository"
	"github.com/WilkerAlves/assistance-go/src/domain/service"
	"github.com/stretchr/testify/mock"
)

type MyMockedCategoryService struct {
	mock.Mock
	Repo repository.ICategoryRepository
}

func (s *MyMockedCategoryService) Create(category entity.Category) error {
	if cat, _ := s.Repo.FindByName(category.GetName()); cat != nil {
		return errors.New("the category name already exists")
	}

	err := s.Repo.Create(category)
	if err != nil {
		return fmt.Errorf("error while create category, %w", err)
	}
	return nil
}
func (s *MyMockedCategoryService) Update(category entity.Category) error {
	return nil
}
func (s *MyMockedCategoryService) GetById(id string) (*entity.Category, error) {
	return nil, nil
}
func (s *MyMockedCategoryService) GetByName(name string) (*entity.Category, error) {
	return nil, nil
}
func (s *MyMockedCategoryService) GetAll(filters *service.CategoryFiltersDTO) ([]*entity.Category, error) {
	return s.Repo.FindAll(filters.Active)
}
