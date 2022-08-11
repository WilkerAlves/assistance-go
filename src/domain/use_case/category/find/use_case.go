package find

import "github.com/WilkerAlves/assistance-go/src/domain/service"

type FindCategoryUseCase struct {
	categoryService service.ICategoryService
}

func (f *FindCategoryUseCase) Execute() ([]OutputCategory, error) {
	categories, err := f.categoryService.GetAll()
	if err != nil {
		return nil, err
	}

	output := make([]OutputCategory, 0)
	for _, category := range categories {
		output = append(output, OutputCategory{
			Name:          category.GetName(),
			SupplierTotal: len(category.GetSuppliers()),
		})
	}

	return output, nil
}

func NewFindCategoryUseCase(service service.ICategoryService) *FindCategoryUseCase {
	return &FindCategoryUseCase{categoryService: service}
}
