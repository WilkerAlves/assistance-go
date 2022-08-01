package entity_test

import (
	"testing"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateCategory(t *testing.T) {
	id := uuid.New().String()
	category, err := entity.NewCategory("Categoria1", "sale", &id)
	assert.Nil(t, err)
	assert.Equal(t, id, category.GetID())
	assert.Equal(t, "Categoria1", category.GetName())
	assert.Equal(t, "sale", category.GetAssistanceType())
	assert.Equal(t, 0, len(category.GetSubcategories()))
}

func TestShouldCreateCategoryWhenNotID(t *testing.T) {
	category, err := entity.NewCategory("Categoria1", "sale", nil)
	assert.Nil(t, err)
	assert.Equal(t, "", category.GetID())
	assert.Equal(t, "Categoria1", category.GetName())
	assert.Equal(t, "sale", category.GetAssistanceType())
	assert.Equal(t, 0, len(category.GetSubcategories()))
}

func TestShouldReturnErrorWhenCreateCategoryNameEmpty(t *testing.T) {
	id := uuid.New().String()

	_, err := entity.NewCategory("", "sale", &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")

	_, err = entity.NewCategory(" ", "sale", &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")
}

func TestShouldReturnErrorWhenCreateCategoryAssistanceTypeEmpty(t *testing.T) {
	id := uuid.New().String()

	_, err := entity.NewCategory("Categoria1", "", &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")

	_, err = entity.NewCategory("Categoria1", " ", &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")

	_, err = entity.NewCategory("Categoria1", "xxxxxxx", &id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the assistance type is invalid")
}

func TestShouldAddSubcategory(t *testing.T) {
	id := uuid.New().String()
	category, _ := entity.NewCategory("Categoria1", "sale", &id)

	subcategory, _ := entity.NewCategory("Subcategoria1", "paid", &id)
	err := category.AddSubcategory(*subcategory)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(category.GetSubcategories()))
}

func TestShouldReturnErrorWhenAddSubcategoryWithNameAlreadyExists(t *testing.T) {
	id := uuid.New().String()
	category, _ := entity.NewCategory("Categoria1", "sale", &id)

	subcategory, _ := entity.NewCategory("Subcategoria1", "paid", &id)
	err := category.AddSubcategory(*subcategory)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(category.GetSubcategories()))

	subcategory2, _ := entity.NewCategory("Subcategoria1", "paid", &id)
	err = category.AddSubcategory(*subcategory2)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "already exists a subcategory with this name")
	assert.Equal(t, 1, len(category.GetSubcategories()))
}

func TestShouldUpdateNameCategory(t *testing.T) {
	oldName := "Categoria1"
	newName := "NewNameCategoria"
	category, err := entity.NewCategory(oldName, "sale", nil)
	assert.Nil(t, err)

	err = category.ChangeName(newName)

	assert.Nil(t, err)
	assert.Equal(t, newName, category.GetName())
}

func TestShouldUpdateAssistanceType(t *testing.T) {
	oldType := "sale"
	newType := "paid"
	category, err := entity.NewCategory("Categoria1", oldType, nil)
	assert.Nil(t, err)

	err = category.ChangeAssistanceType(newType)

	assert.Nil(t, err)
	assert.Equal(t, newType, category.GetAssistanceType())
}

func TestShouldReturnErrorWhenUpdateNameCategoryInvalid(t *testing.T) {
	oldName := "Categoria1"
	category, err := entity.NewCategory(oldName, "sale", nil)
	assert.Nil(t, err)

	err = category.ChangeName(" ")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")

	err = category.ChangeName("")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the category name is empty")
}

func TestShouldReturnErrorWhenUpdateAssistanceTypeInvalid(t *testing.T) {
	category, err := entity.NewCategory("Categoria1", "sale", nil)
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
