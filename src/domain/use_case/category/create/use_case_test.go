package create_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/interface/repository"
	"github.com/WilkerAlves/assistance-go/src/domain/mocks"
	"github.com/WilkerAlves/assistance-go/src/domain/use_case/category/create"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyMockedEventService struct {
	mock.Mock
}

func (m *MyMockedEventService) Send(eventName string, body interface{}) bool {
	return true
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
	return nil, nil
}

type MyMockedGeneratedIdsService struct {
	mock.Mock
}

func (m *MyMockedGeneratedIdsService) Create() (string, error) {
	return "12345677", nil
}

func TestCreateCategoryUseCase_Execute(t *testing.T) {
	repositoryMock := new(mocks.MyMockedCategoryRepository)
	eventServiceMock := new(MyMockedEventService)
	categoryServiceMock := new(MyMockedCategoryService)
	generatedIdsServiceMock := new(MyMockedGeneratedIdsService)
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
	eventServiceMock := new(MyMockedEventService)
	categoryServiceMock := new(MyMockedCategoryService)
	generatedIdsServiceMock := new(MyMockedGeneratedIdsService)
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
