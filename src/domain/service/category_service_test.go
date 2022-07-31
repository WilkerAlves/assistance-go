package service_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/service"
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

func TestShouldCreateCategory(t *testing.T) {
	repositoryMock := new(MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	category, _ := entity.NewCategory("CategoryName", "sale", nil)

	err := categoryService.Create(*category)

	assert.Nil(t, err)
}

func TestShouldReturnErroWhenCreateCategoryWithNameAlreadyExists(t *testing.T) {
	repositoryMock := new(MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	category, _ := entity.NewCategory("CategoryName", "sale", nil)

	err := categoryService.Create(*category)
	assert.Nil(t, err)

	category2, _ := entity.NewCategory("CategoryName", "sale", nil)
	err = categoryService.Create(*category2)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name already exists")
}

func TestShouldReturnCategoryEntityWhenSearchByName(t *testing.T) {
	repositoryMock := new(MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	name := "CategoryName1"

	category, _ := entity.NewCategory(name, "sale", nil)
	_ = categoryService.Create(*category)

	cat, err := categoryService.GetByName(name)

	assert.Nil(t, err)
	assert.Equal(t, name, cat.GetName())
	assert.Equal(t, "sale", cat.GetAssistanceType())
	assert.Equal(t, "", cat.GetID())
}

func TestShouldReturnCategoryEntityWhenSearchById(t *testing.T) {
	repositoryMock := new(MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	name := "CategoryName1"
	id := uuid.New().String()

	category, _ := entity.NewCategory(name, "sale", &id)
	_ = categoryService.Create(*category)

	cat, err := categoryService.GetById(id)

	assert.Nil(t, err)
	assert.Equal(t, name, cat.GetName())
	assert.Equal(t, "sale", cat.GetAssistanceType())
	assert.Equal(t, id, cat.GetID())
}

func TestShouldUpdateCategory(t *testing.T) {
	repositoryMock := new(MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	id := uuid.New().String()
	name := "CategoryName1"
	category, _ := entity.NewCategory(name, "sale", &id)
	_ = categoryService.Create(*category)
	cat, err := categoryService.GetByName(name)

	err = cat.ChangeName("NewCategoryName1")
	assert.Nil(t, err)

	err = cat.ChangeAssistanceType("paid")
	assert.Nil(t, err)

	err = categoryService.Update(*cat)
	assert.Nil(t, err)

	newCat, err := categoryService.GetById(cat.GetID())

	assert.Equal(t, "NewCategoryName1", newCat.GetName())
	assert.Equal(t, "paid", newCat.GetAssistanceType())
	assert.Equal(t, id, newCat.GetID())
}
