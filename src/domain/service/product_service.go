package service

import "github.com/WilkerAlves/assistance-go/src/domain/interface/repository"

type IProductService interface {
	FindByCategoryID(categoryID string) ([]string, error)
}

type productService struct {
	repo repository.IProductRepository
}

func (p *productService) FindByCategoryId(categoryID string) ([]string, error) {
	products, err := p.repo.FindByCategoryID(categoryID)
	if err != nil {
		return make([]string, 0), err
	}

	return products, nil
}
