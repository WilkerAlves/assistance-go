package create_test

import (
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/mocks"
	"github.com/WilkerAlves/assistance-go/src/domain/use_case/category/create"
	"github.com/stretchr/testify/suite"
)

type CreateCategoryUseCaseTestSuite struct {
	suite.Suite
	eventService    *mocks.MyMockedEventService
	categoryService *mocks.MyMockedCategoryService
	generateIds     *mocks.MyMockedGeneratedIdsService
}

func (c *CreateCategoryUseCaseTestSuite) SetupTest() {
	c.eventService = new(mocks.MyMockedEventService)
	c.generateIds = new(mocks.MyMockedGeneratedIdsService)
	c.categoryService = new(mocks.MyMockedCategoryService)

	repositoryMock := new(mocks.MyMockedCategoryRepository)
	c.categoryService.Repo = repositoryMock
}

func (c *CreateCategoryUseCaseTestSuite) TestShouldNotReturnErrorWhenCreateCreateWithAllFieldsValid() {
	useCase := create.NewCreateCategoryUseCase(
		c.eventService,
		c.categoryService,
		c.generateIds,
	)

	input := create.InputCrateCategory{
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
