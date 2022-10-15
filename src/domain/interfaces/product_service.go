package interfaces

type IProductService interface {
	FindByCategoryID(categoryID string) ([]string, error)
}
