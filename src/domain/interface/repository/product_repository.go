package repository

type IProductRepository interface {
	FindByCategoryID(categoryID string) ([]string, error)
}
