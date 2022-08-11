package find_test

import (
	"errors"
	"fmt"
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
	return nil
}
func (m *MyMockedCategoryRepository) Find(id string) (*entity.Category, error) {
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
func (m *MyMockedCategoryRepository) FindAll() ([]*entity.Category, error) {
	output := make([]*entity.Category, 0)
	for _, category := range m.DB {
		output = append(output, &category)
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
func (s *MyMockedCategoryService) GetAll() ([]*entity.Category, error) {
	return s.Repo.FindAll()
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
	outputCategories, err := useCase.Execute()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(outputCategories))
}

func TestCreateCategoryUseCase_Execute(t *testing.T) {
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
	outputCategories, err := useCase.Execute()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(outputCategories))
	assert.Equal(t, name, outputCategories[0].Name)
	assert.Equal(t, 2, outputCategories[0].SupplierTotal)
}
