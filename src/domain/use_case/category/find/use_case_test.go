package find_test

import (
	"errors"
	"fmt"
	"github.com/WilkerAlves/assistance-go/src/domain/service"
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/interface/repository"
	"github.com/WilkerAlves/assistance-go/src/domain/use_case/category/find"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyMockedCategoryRepository struct {
	mock.Mock
	DB []entity.Category
}

func (m *MyMockedCategoryRepository) Create(category entity.Category) error {
	m.DB = append(m.DB, category)
	return nil
}

func (m *MyMockedCategoryRepository) Update(category entity.Category) error {
	for i, oldCategory := range m.DB {
		if oldCategory.GetID() == category.GetID() {
			err := oldCategory.ChangeName(category.GetName())
			if err != nil {
				return err
			}

			err = oldCategory.ChangeAssistanceType(category.GetAssistanceType())
			if err != nil {
				return err
			}

			if category.GetStatus() == false {
				oldCategory.Inactivate()
			}

			m.DB[i] = oldCategory

			return nil
		}
	}
	return errors.New(fmt.Sprintf("category not found. id: %s", category.GetID()))
}

func (m *MyMockedCategoryRepository) Find(id string) (*entity.Category, error) {
	for _, category := range m.DB {
		if category.GetID() == id {
			return &category, nil
		}
	}
	return nil, nil
}

func (m *MyMockedCategoryRepository) FindByName(name string) (*entity.Category, error) {
	for _, category := range m.DB {
		if category.GetName() == name {
			return &category, nil
		}
	}
	return nil, nil
}

func (m *MyMockedCategoryRepository) FindAll(active *bool) ([]*entity.Category, error) {
	output := make([]*entity.Category, 0)
	if active == nil {
		for _, category := range m.DB {
			output = append(output, &category)
		}

		return output, nil
	}

	for i := range m.DB {
		if m.DB[i].GetStatus() == *active {
			output = append(output, &m.DB[i])
		}
	}

	return output, nil
}

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

func TestShouldReturnListOutputCategory(t *testing.T) {
	repositoryMock := new(MyMockedCategoryRepository)
	categoryServiceMock := new(MyMockedCategoryService)
	categoryServiceMock.Repo = repositoryMock

	id := uuid.New().String()
	id2 := uuid.New().String()
	name := "CategoryName1"
	name2 := "CategoryName2"
	category, _ := entity.NewCategory(name, "sale", "1234", &id)
	category2, _ := entity.NewCategory(name2, "paid", "56789", &id2)

	_ = categoryServiceMock.Repo.Create(*category)
	_ = categoryServiceMock.Repo.Create(*category2)

	useCase := find.NewFindCategoryUseCase(categoryServiceMock)
	inputFilterCategory := find.InputFilterCategory{}
	outputCategories, err := useCase.Execute(inputFilterCategory)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(outputCategories))
}

func TestShouldFindCategoryUseCase_Execute(t *testing.T) {
	repositoryMock := new(MyMockedCategoryRepository)
	categoryServiceMock := new(MyMockedCategoryService)
	categoryServiceMock.Repo = repositoryMock

	id := uuid.New().String()
	name := "CategoryName1"
	category, _ := entity.NewCategory(name, "sale", "1234", &id)

	_ = categoryServiceMock.Repo.Create(*category)

	_ = category.AddSupplier("001756")
	_ = category.AddSupplier("001757")

	useCase := find.NewFindCategoryUseCase(categoryServiceMock)
	inputFilterCategory := find.InputFilterCategory{}
	outputCategories, err := useCase.Execute(inputFilterCategory)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(outputCategories))
	assert.Equal(t, name, outputCategories[0].Name)
	assert.Equal(t, 2, outputCategories[0].SupplierTotal)
}

func TestShouldReturnListActiveCategoriesFindUseCase(t *testing.T) {
	repositoryMock := new(MyMockedCategoryRepository)
	categoryServiceMock := new(MyMockedCategoryService)
	categoryServiceMock.Repo = repositoryMock

	id := uuid.New().String()
	id2 := uuid.New().String()
	name := "CategoryName1"
	name2 := "CategoryName2"
	category, _ := entity.NewCategory(name, "sale", "1234", &id)
	category2, _ := entity.NewCategory(name2, "sale", "1234", &id2)

	_ = categoryServiceMock.Repo.Create(*category)
	_ = categoryServiceMock.Repo.Create(*category2)

	category2.Inactivate()
	_ = categoryServiceMock.Repo.Update(*category2)

	useCase := find.NewFindCategoryUseCase(categoryServiceMock)
	active := true
	inputFilterCategory := find.InputFilterCategory{Active: &active}
	outputCategories, err := useCase.Execute(inputFilterCategory)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(outputCategories))
	assert.Equal(t, name, outputCategories[0].Name)
}
