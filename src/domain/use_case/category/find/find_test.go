package find_test

import (
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/dto"
	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/mocks"
	"github.com/WilkerAlves/assistance-go/src/domain/use_case/category/find"
	"github.com/stretchr/testify/suite"
)

type FindCategoryUseCaseTestSuite struct {
	suite.Suite
	categoryService *mocks.MyMockedCategoryService
}

func (f *FindCategoryUseCaseTestSuite) SetupTest() {
	f.categoryService = new(mocks.MyMockedCategoryService)

	repositoryMock := new(mocks.MyMockedCategoryRepository)
	f.categoryService.Repo = repositoryMock

	id := "1"
	category1, _ := entity.NewCategory("CategoryName1", "sale", "1234", &id)
	_ = category1.AddSupplier("001756")
	_ = category1.AddSupplier("001757")

	id2 := "2"
	category2, _ := entity.NewCategory("CategoryName2", "paid", "5678", &id2)

	id3 := "3"
	category3, _ := entity.NewCategory("CategoryName3", "subsidized", "9123", &id3)

	_ = f.categoryService.Repo.Create(*category1)
	_ = f.categoryService.Repo.Create(*category2)
	_ = f.categoryService.Repo.Create(*category3)
}

func (f *FindCategoryUseCaseTestSuite) TearDownTestSuite() {
	f.categoryService = nil
}

func (f *FindCategoryUseCaseTestSuite) TestShouldReturnListOutputCategory() {
	useCase := find.NewFindCategoryUseCase(f.categoryService)
	outputCategories, err := useCase.Execute(dto.InputFilterCategory{})
	f.Assert().Nil(err)
	f.Assert().Equal(3, len(outputCategories))
}

func (f *FindCategoryUseCaseTestSuite) TestShouldFindCategoryUseCase() {
	useCase := find.NewFindCategoryUseCase(f.categoryService)
	outputCategories, err := useCase.Execute(dto.InputFilterCategory{})

	var output dto.OutputCategory

	for i := range outputCategories {
		if outputCategories[i].Name == "CategoryName1" {
			output = outputCategories[i]
		}

	}

	f.Assert().Nil(err)
	f.Assert().Equal(3, len(outputCategories))
	f.Assert().Equal("CategoryName1", output.Name)
	f.Assert().Equal(2, output.SupplierTotal)
}

func (f *FindCategoryUseCaseTestSuite) TestShouldReturnListActiveCategoriesFindUseCase() {
	category, _ := f.categoryService.Repo.Find("2")
	category.Inactivate()
	_ = f.categoryService.Repo.Update(*category)

	useCase := find.NewFindCategoryUseCase(f.categoryService)

	active := true
	outputCategories, err := useCase.Execute(dto.InputFilterCategory{Active: &active})

	f.Assert().Nil(err)
	f.Assert().Equal(2, len(outputCategories))
}

func TestSuiteFindCategoryUseCase(t *testing.T) {
	suite.Run(t, &FindCategoryUseCaseTestSuite{})
}
