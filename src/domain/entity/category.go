package entity

import (
	"errors"
	"strings"
)

const (
	Sale                            = "sale"
	Paid                            = "paid"
	Subsidized                      = "subsidized"
	ValidateMessageToName           = "the category name is empty"
	ValidateMessageToAssistanceType = "the assistance type is invalid"
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
	if !validateNameCategory(newName) {
		return errors.New(ValidateMessageToName)
	}
	c.name = newName
	return nil
}

func (c *Category) ChangeAssistanceType(newType string) error {
	if !validateAssistanceType(newType) {
		return errors.New(ValidateMessageToAssistanceType)
	}
	c.assistanceType = newType
	return nil
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

func NewCategory(name, assistanceType string, id *string) (*Category, error) {
	category := &Category{name: name, assistanceType: assistanceType, subcategories: make([]Category, 0)}

	if !validateNameCategory(category.name) {
		return nil, errors.New(ValidateMessageToName)
	}

	if !validateAssistanceType(category.assistanceType) {
		return nil, errors.New(ValidateMessageToAssistanceType)
	}

	if id != nil {
		category.id = *id
	}

	return category, nil
}

func validateNameCategory(name string) bool {
	if len(strings.Trim(name, " ")) < 1 {
		return false
	}
	return true
}

func validateAssistanceType(assistanceType string) bool {
	switch assistanceType {
	case Sale:
		return true
	case Subsidized:
		return true
	case Paid:
		return true
	}
	return false
}
