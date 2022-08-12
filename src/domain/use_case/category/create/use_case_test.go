package create_test

import (
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/mocks"
	"github.com/WilkerAlves/assistance-go/src/domain/use_case/category/create"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategoryUseCase_Execute(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	eventServiceMock := new(mocks.MyMockedEventService)
	categoryServiceMock := new(mocks.MyMockedCategoryService)
	generatedIdsServiceMock := new(mocks.MyMockedGeneratedIdsService)
	categoryServiceMock.Repo = repositoryMock

	useCase := new(create.CreateCategoryUseCase)
	useCase.CategoryService = categoryServiceMock
	useCase.EventService = eventServiceMock
	useCase.GenerateIds = generatedIdsServiceMock

	input := create.InputCrateCategory{
		Name:           "CategoryUseCase",
		AssistanceType: "sale",
	}

	err := useCase.Execute(input)

	assert.Nil(t, err)
}

func TestCreateCategoryUseCase_Execute_ShouldReturnErrorWhenCategoryNameInvalid(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	eventServiceMock := new(mocks.MyMockedEventService)
	categoryServiceMock := new(mocks.MyMockedCategoryService)
	generatedIdsServiceMock := new(mocks.MyMockedGeneratedIdsService)
	categoryServiceMock.Repo = repositoryMock

	useCase := new(create.CreateCategoryUseCase)
	useCase.CategoryService = categoryServiceMock
	useCase.EventService = eventServiceMock
	useCase.GenerateIds = generatedIdsServiceMock

	inputCorrect := create.InputCrateCategory{
		Name:           "CategoryUseCase",
		AssistanceType: "sale",
	}
	inputNameInvalid := create.InputCrateCategory{
		Name:           "",
		AssistanceType: "sale",
	}
	inputNameAlreadyExists := create.InputCrateCategory{
		Name:           "CategoryUseCase",
		AssistanceType: "sale",
	}

	err := useCase.Execute(inputCorrect)
	assert.Nil(t, err)

	err = useCase.Execute(inputNameInvalid)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")

	err = useCase.Execute(inputNameAlreadyExists)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name already exists")
}
