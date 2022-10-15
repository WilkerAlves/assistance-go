package use_case_test

import (
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/dto"
	"github.com/WilkerAlves/assistance-go/src/domain/use_case/category/create"
	mocks2 "github.com/WilkerAlves/assistance-go/tests/mocks"
	"github.com/stretchr/testify/suite"
)

type CreateCategoryUseCaseTestSuite struct {
	suite.Suite
	eventService    *mocks2.MyMockedEventService
	categoryService *mocks2.MyMockedCategoryService
	generateIds     *mocks2.MyMockedGeneratedIdsService
}

func (c *CreateCategoryUseCaseTestSuite) SetupTest() {
	c.eventService = new(mocks2.MyMockedEventService)
	c.generateIds = new(mocks2.MyMockedGeneratedIdsService)
	c.categoryService = new(mocks2.MyMockedCategoryService)

	repositoryMock := new(mocks2.MyMockedCategoryRepository)
	c.categoryService.Repo = repositoryMock
}

func (c *CreateCategoryUseCaseTestSuite) TestShouldNotReturnErrorWhenCreateCreateWithAllFieldsValid() {
	useCase := create.NewCreateCategoryUseCase(
		c.eventService,
		c.categoryService,
		c.generateIds,
	)

	input := dto.InputCrateCategory{
		Name:           "CategoryUseCase",
		AssistanceType: "sale",
	}

	err := useCase.Execute(input)

	c.Assert().Nil(err)
}

func (c *CreateCategoryUseCaseTestSuite) TestShouldReturnErrorWhenAttemptCreateCategoryWithNameInvalid() {
	useCase := create.NewCreateCategoryUseCase(
		c.eventService,
		c.categoryService,
		c.generateIds,
	)

	inputCorrect := dto.InputCrateCategory{
		Name:           "CategoryUseCase",
		AssistanceType: "sale",
	}
	inputNameInvalid := dto.InputCrateCategory{
		Name:           "",
		AssistanceType: "sale",
	}
	inputNameAlreadyExists := dto.InputCrateCategory{
		Name:           "CategoryUseCase",
		AssistanceType: "sale",
	}

	err := useCase.Execute(inputCorrect)
	c.Assert().Nil(err)

	err = useCase.Execute(inputNameInvalid)
	c.Assert().NotNil(err)
	c.Assert().EqualError(err, "the category name is empty")

	err = useCase.Execute(inputNameAlreadyExists)
	c.Assert().NotNil(err)
	c.Assert().EqualError(err, "the category name already exists")

}

func TestSuiteCreateCategoryUseCase(t *testing.T) {
	suite.Run(t, &CreateCategoryUseCaseTestSuite{})
}
