package service_test

import (
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/mocks"
	"github.com/WilkerAlves/assistance-go/src/domain/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateCategory(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	category, _ := entity.NewCategory("CategoryName", "sale", "1234", nil)

	err := categoryService.Create(*category)

	assert.Nil(t, err)
}

func TestShouldReturnErrorWhenCreateCategoryWithNameAlreadyExists(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	category, _ := entity.NewCategory("CategoryName", "sale", "1234", nil)

	err := categoryService.Create(*category)
	assert.Nil(t, err)

	category2, _ := entity.NewCategory("CategoryName", "sale", "1234", nil)
	err = categoryService.Create(*category2)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name already exists")
}

func TestShouldReturnCategoryEntityWhenSearchByName(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	name := "CategoryName1"

	category, _ := entity.NewCategory(name, "sale", "1234", nil)
	_ = categoryService.Create(*category)

	cat, err := categoryService.GetByName(name)

	assert.Nil(t, err)
	assert.Equal(t, name, cat.GetName())
	assert.Equal(t, "sale", cat.GetAssistanceType())
	assert.Equal(t, "", cat.GetID())
}

func TestShouldReturnCategoryEntityWhenSearchById(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	name := "CategoryName1"
	id := uuid.New().String()

	category, _ := entity.NewCategory(name, "sale", "1234", &id)
	_ = categoryService.Create(*category)

	cat, err := categoryService.GetById(id)

	assert.Nil(t, err)
	assert.Equal(t, name, cat.GetName())
	assert.Equal(t, "sale", cat.GetAssistanceType())
	assert.Equal(t, id, cat.GetID())
}

func TestShouldUpdateCategory(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)
	id := uuid.New().String()
	name := "CategoryName1"
	category, _ := entity.NewCategory(name, "sale", "1234", &id)
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

func TestShouldReturnListCategories(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	repositoryMock.DB = make([]entity.Category, 0)
	categoryService := service.NewCategoryService(repositoryMock)

	id := uuid.New().String()
	id2 := uuid.New().String()
	name := "CategoryName1"
	name2 := "CategoryName2"
	category, _ := entity.NewCategory(name, "sale", "1234", &id)
	category2, _ := entity.NewCategory(name2, "paid", "56789", &id2)
	_ = categoryService.Create(*category)
	_ = categoryService.Create(*category2)

	categories, err := categoryService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(categories))
}
