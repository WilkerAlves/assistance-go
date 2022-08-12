package find_test

import (
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/mocks"
	"github.com/WilkerAlves/assistance-go/src/domain/use_case/category/find"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnListOutputCategory(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	categoryServiceMock := new(mocks.MyMockedCategoryService)
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
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	categoryServiceMock := new(mocks.MyMockedCategoryService)
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
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	categoryServiceMock := new(mocks.MyMockedCategoryService)
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
