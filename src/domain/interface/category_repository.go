package _interface

import "github.com/WilkerAlves/assistance-go/src/domain/entity"

type ICategoryRepository interface {
	Create(category entity.Category) error
	Update(category entity.Category) error
	Find(id string) (*entity.Category, error)
	FindByName(name string) (*entity.Category, error)
}
