package service

import (
	"github.com/WilkerAlves/assistance-go/src/domain/interfaces"
)

type IProductService interface {
	FindByCategoryID(categoryID string) ([]string, error)
}

type productService struct {
	repo interfaces.IProductRepository
}

func (p *productService) FindByCategoryId(categoryID string) ([]string, error) {
	products, err := p.repo.FindByCategoryID(categoryID)
	if err != nil {
		return make([]string, 0), err
	}

	return products, nil
}
