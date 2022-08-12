package dto

type CategoryFiltersDTO struct {
	Active *bool
}

func NewCategoryFiltersDTO(active *bool) *CategoryFiltersDTO {
	return &CategoryFiltersDTO{
		Active: active,
	}
}
