package entity

import (
	"errors"
	"strings"
)

const (
	SALE       = "sale"
	PAID       = "paid"
	SUBSIDIZED = "subsidized"
)

type Category struct {
	id             string
	name           string
	assistanceType string
	subcategories  []Category
}

func (c *Category) GetID() string {
	return c.id
}

func (c *Category) GetName() string {
	return c.name
}

func (c *Category) ChangeName(newName string) error {
	c.name = newName
	return c.Valid()
}

func (c *Category) ChangeAssistanceType(newType string) error {
	c.assistanceType = newType
	return c.Valid()
}

func (c *Category) GetAssistanceType() string {
	return c.assistanceType
}

func (c *Category) GetSubcategories() []Category {
	return c.subcategories
}

func (c *Category) AddSubcategory(category Category) error {
	for _, subcategory := range c.subcategories {
		if subcategory.name == category.name {
			return errors.New("already exists a subcategory with this name")
		}
	}

	c.subcategories = append(c.subcategories, category)
	return nil
}

func (c *Category) Valid() error {
	if len(strings.Trim(c.name, " ")) < 1 {
		return errors.New("the category name is empty")
	}

	switch c.assistanceType {
	case SALE:
		return nil
	case SUBSIDIZED:
		return nil
	case PAID:
		return nil
	}

	return errors.New("the assistance type is invalid")
}

func NewCategory(name, assistanceType string, id *string) (*Category, error) {
	category := &Category{name: name, assistanceType: assistanceType, subcategories: make([]Category, 0)}

	if err := category.Valid(); err != nil {
		return nil, err
	}

	if id != nil {
		category.id = *id
	}

	return category, nil
}

// Quebrar o metodo de validação em dois para que não deixa a entidade invalida