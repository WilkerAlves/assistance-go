package entity_test

import (
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const StockGroup = "1901"

func TestShouldCreateCategory(t *testing.T) {
	id := uuid.New().String()
	category, err := entity.NewCategory("Categoria1", "sale", StockGroup, &id)
	assert.Nil(t, err)
	assert.Equal(t, id, category.GetID())
	assert.Equal(t, "Categoria1", category.GetName())
	assert.Equal(t, "sale", category.GetAssistanceType())
	assert.Equal(t, 0, len(category.GetSubcategories()))
}

func TestShouldCreateCategoryWhenNotID(t *testing.T) {
	category, err := entity.NewCategory("Categoria1", "sale", StockGroup, nil)
	assert.Nil(t, err)
	assert.Equal(t, "", category.GetID())
	assert.Equal(t, "Categoria1", category.GetName())
	assert.Equal(t, "sale", category.GetAssistanceType())
	assert.Equal(t, 0, len(category.GetSubcategories()))
}

func TestShouldReturnErrorWhenCreateCategoryNameEmpty(t *testing.T) {
	id := uuid.New().String()

	_, err := entity.NewCategory("", "sale", StockGroup, &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")

	_, err = entity.NewCategory(" ", "sale", StockGroup, &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")
}

func TestShouldReturnErrorWhenCreateCategoryAssistanceTypeEmpty(t *testing.T) {
	id := uuid.New().String()

	_, err := entity.NewCategory("Categoria1", "", StockGroup, &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")

	_, err = entity.NewCategory("Categoria1", " ", StockGroup, &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")

	_, err = entity.NewCategory("Categoria1", "xxxxxxx", StockGroup, &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")
}

func TestShouldAddSubcategory(t *testing.T) {
	id := uuid.New().String()
	category, _ := entity.NewCategory("Categoria1", "sale", StockGroup, &id)

	subcategory, _ := entity.NewCategory("Subcategoria1", "paid", StockGroup, &id)
	err := category.AddSubcategory(*subcategory)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(category.GetSubcategories()))
}

func TestShouldReturnErrorWhenAddSubcategoryWithNameAlreadyExists(t *testing.T) {
	id := uuid.New().String()
	category, _ := entity.NewCategory("Categoria1", "sale", StockGroup, &id)

	subcategory, _ := entity.NewCategory("Subcategoria1", "paid", StockGroup, &id)
	err := category.AddSubcategory(*subcategory)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(category.GetSubcategories()))

	subcategory2, _ := entity.NewCategory("Subcategoria1", "paid", StockGroup, &id)
	err = category.AddSubcategory(*subcategory2)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "already exists a subcategory with this name")
	assert.Equal(t, 1, len(category.GetSubcategories()))
}

func TestShouldUpdateNameCategory(t *testing.T) {
	oldName := "Categoria1"
	newName := "NewNameCategoria"
	category, err := entity.NewCategory(oldName, "sale", StockGroup, nil)
	assert.Nil(t, err)

	err = category.ChangeName(newName)

	assert.Nil(t, err)
	assert.Equal(t, newName, category.GetName())
}

func TestShouldUpdateAssistanceType(t *testing.T) {
	oldType := "sale"
	newType := "paid"
	category, err := entity.NewCategory("Categoria1", oldType, StockGroup, nil)
	assert.Nil(t, err)

	err = category.ChangeAssistanceType(newType)

	assert.Nil(t, err)
	assert.Equal(t, newType, category.GetAssistanceType())
}

func TestShouldReturnErrorWhenUpdateNameCategoryInvalid(t *testing.T) {
	oldName := "Categoria1"
	category, err := entity.NewCategory(oldName, "sale", StockGroup, nil)
	assert.Nil(t, err)

	err = category.ChangeName(" ")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")

	err = category.ChangeName("")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")
}

func TestShouldReturnErrorWhenUpdateAssistanceTypeInvalid(t *testing.T) {
	category, err := entity.NewCategory("Categoria1", "sale", StockGroup, nil)
	assert.Nil(t, err)

	err = category.ChangeAssistanceType("xxxxxxxxxx")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")

	err = category.ChangeAssistanceType(" ")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")

	err = category.ChangeAssistanceType("")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")
}

func TestShouldBindBetweenSupplierIdAndCategory(t *testing.T) {
	category, _ := entity.NewCategory("Categoria1", "sale", StockGroup, nil)
	assert.Equal(t, len(category.GetSubcategories()), 0)

	supplierId := "001756"

	err := category.AddSupplier(supplierId)
	assert.Nil(t, err)

	assert.Equal(t, len(category.GetSuppliers()), 1)
}

func TestShouldBindBetweenSupplierIdAndCategories(t *testing.T) {
	category, _ := entity.NewCategory("Categoria1", "sale", StockGroup, nil)
	assert.Equal(t, len(category.GetSubcategories()), 0)

	category2, _ := entity.NewCategory("Categoria2", "sale", StockGroup, nil)
	assert.Equal(t, len(category2.GetSubcategories()), 0)

	supplierId := "001756"

	err := category.AddSupplier(supplierId)
	assert.Nil(t, err)

	err = category2.AddSupplier(supplierId)
	assert.Nil(t, err)

	assert.Equal(t, len(category.GetSuppliers()), 1)
	assert.Equal(t, len(category2.GetSuppliers()), 1)
}

func TestShouldRemoveSubcategory(t *testing.T) {
	categoryID := uuid.New().String()
	sub1ID := uuid.New().String()
	sub2ID := uuid.New().String()

	category, _ := entity.NewCategory("Categoria1", "sale", StockGroup, &categoryID)

	subcategory, _ := entity.NewCategory("Subcategoria1", "paid", StockGroup, &sub1ID)
	err := category.AddSubcategory(*subcategory)

	subcategory2, _ := entity.NewCategory("Subcategoria2", "paid", StockGroup, &sub2ID)
	err = category.AddSubcategory(*subcategory2)

	assert.Equal(t, 2, len(category.GetSubcategories()))

	err = category.RemoveSubcategory(*subcategory2)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(category.GetSubcategories()))

	res, err := category.GetSubcategory(sub1ID)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = category.GetSubcategory(sub2ID)
	assert.Nil(t, err)
	assert.Nil(t, res)
}

func TestShouldInactivateCategoryAndSubcategories(t *testing.T) {
	categoryID := uuid.New().String()
	sub1ID := uuid.New().String()
	sub2ID := uuid.New().String()

	category, _ := entity.NewCategory("Categoria1", "sale", StockGroup, &categoryID)

	subcategory, _ := entity.NewCategory("Subcategoria1", "paid", StockGroup, &sub1ID)
	err := category.AddSubcategory(*subcategory)
	assert.Nil(t, err)

	subcategory2, _ := entity.NewCategory("Subcategoria2", "paid", StockGroup, &sub2ID)
	err = category.AddSubcategory(*subcategory2)
	assert.Nil(t, err)

	category.Inactivate()

	sub, _ := category.GetSubcategory(sub1ID)
	sub2, _ := category.GetSubcategory(sub2ID)

	assert.Equal(t, false, category.GetStatus())
	assert.Equal(t, false, sub.GetStatus())
	assert.Equal(t, false, sub2.GetStatus())
}

func TestShouldInactivateSubcategory(t *testing.T) {
	categoryID := uuid.New().String()
	sub1ID := uuid.New().String()

	category, _ := entity.NewCategory("Categoria1", "sale", StockGroup, &categoryID)

	subcategory, _ := entity.NewCategory("Subcategoria1", "paid", StockGroup, &sub1ID)
	err := category.AddSubcategory(*subcategory)
	assert.Nil(t, err)

	err = category.InactivateSubCategory(sub1ID)
	assert.Nil(t, err)

	sub, _ := category.GetSubcategory(sub1ID)

	assert.Equal(t, true, category.GetStatus())
	assert.Equal(t, false, sub.GetStatus())
}

func TestShouldReturnErrorWhenCreateCategoryWithStockGroupInvalid(t *testing.T) {
	_, err := entity.NewCategory("Categoria1", "sale", "", nil)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the stock group is empty")

	_, err = entity.NewCategory("Categoria1", "sale", " ", nil)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the stock group is empty")
}

func TestShouldChangedStockGroupSubcategory(t *testing.T) {
	categoryID := uuid.New().String()
	category2ID := uuid.New().String()

	category, _ := entity.NewCategory("Categoria1", "sale", StockGroup, &categoryID)
	category2, _ := entity.NewCategory("Categoria2", "paid", StockGroup, &category2ID)

	_ = category.AddSubcategory(*category2)

	subcategory, _ := category.GetSubcategory(category2ID)
	assert.Equal(t, StockGroup, subcategory.GetStockGroup())

	err := category.ChangeStockGroupSubCategory(category2ID, "1234")
	assert.Nil(t, err)

	subcategory, _ = category.GetSubcategory(category2ID)
	assert.Equal(t, "1234", subcategory.GetStockGroup())
}
