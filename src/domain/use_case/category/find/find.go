package find

import (
	"github.com/WilkerAlves/assistance-go/src/domain/dto"
	"github.com/WilkerAlves/assistance-go/src/domain/interfaces"
)

type FindCategoryUseCase struct {
	categoryService interfaces.ICategoryService
}

func (f *FindCategoryUseCase) Execute(inputFilterCategory dto.InputFilterCategory) ([]dto.OutputCategory, error) {
	categoryFilter := dto.NewCategoryFiltersDTO(inputFilterCategory.Active)

	categories, err := f.categoryService.GetAll(categoryFilter)
	if err != nil {
		return nil, err
	}

	output := make([]dto.OutputCategory, 0)
	for _, category := range categories {
		output = append(output, dto.OutputCategory{
			Name:          category.GetName(),
			SupplierTotal: len(category.GetSuppliers()),
		})
	}

	return output, nil
}

func NewFindCategoryUseCase(service interfaces.ICategoryService) *FindCategoryUseCase {
	return &FindCategoryUseCase{categoryService: service}
}
