package interfaces

type IProductRepository interface {
	FindByCategoryID(categoryID string) ([]string, error)
}
