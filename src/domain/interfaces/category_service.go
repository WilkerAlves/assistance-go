package interfaces

import (
	"github.com/WilkerAlves/assistance-go/src/domain/dto"
	"github.com/WilkerAlves/assistance-go/src/domain/entity"
)

type ICategoryService interface {
	Create(category entity.Category) error
	Update(category entity.Category) error
	GetById(id string) (*entity.Category, error)
	GetByName(name string) (*entity.Category, error)
	GetAll(filters *dto.CategoryFiltersDTO) ([]entity.Category, error)
}
