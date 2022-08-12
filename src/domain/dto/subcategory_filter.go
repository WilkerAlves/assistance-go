package dto

type SubCategoryFiltersDTO struct {
	Active *bool
}

func NewSubCategoryFiltersDTO(active *bool) *SubCategoryFiltersDTO {
	return &SubCategoryFiltersDTO{
		Active: active,
	}
}
