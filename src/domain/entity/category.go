package entity

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Sale                            = "sale"
	Paid                            = "paid"
	Subsidized                      = "subsidized"
	ValidateMessageToName           = "the category name is empty"
	ValidateMessageToAssistanceType = "the assistance type is invalid"
	ValidateMessageToSupplierId     = "the supplierId is empty"
)

type Category struct {
	id             string
	name           string
	assistanceType string
	subcategories  map[string]*Subcategory
	suppliers      map[string]string
	active         bool
}

func (c *Category) GetID() string {
	return c.id
}

func (c *Category) GetName() string {
	return c.name
}

func (c *Category) GetStatus() bool {
	return c.active
}

func (c *Category) GetAssistanceType() string {
	return c.assistanceType
}

func (c *Category) ChangeName(newName string) error {
	if !validateStringEmpty(newName) {
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

func (c *Category) Inactivate() {
	c.active = false
	for i := range c.subcategories {
		subcategory := c.subcategories[i]
		subcategory.Inactivate()
		fmt.Println(subcategory.GetStatus())
	}
}

// Subcategories

func (c *Category) GetSubcategories() map[string]*Subcategory {
	return c.subcategories
}

func (c *Category) AddSubcategory(cat Category) error {
	for _, subcategory := range c.subcategories {
		if subcategory.category.GetName() == cat.GetName() {
			return errors.New("already exists a subcategory with this name")
		}
	}
	c.subcategories[cat.GetID()] = &Subcategory{
		category: &cat,
		active:   cat.GetStatus(),
	}
	return nil
}

func (c *Category) RemoveSubcategory(cat Category) error {
	if !validateStringEmpty(cat.GetID()) {
		return errors.New(ValidateMessageToName)
	}

	id := cat.GetID()

	delete(c.subcategories, id)
	return nil
}

func (c *Category) GetSubcategory(subID string) (*Subcategory, error) {
	if !validateStringEmpty(subID) {
		return nil, errors.New(ValidateMessageToName)
	}

	sub := c.subcategories[subID]

	if sub == nil {
		return nil, nil
	}

	return sub, nil
}

func (c *Category) InactivateSubCategory(subID string) error {
	subcategory, err := c.GetSubcategory(subID)
	if err != nil {
		return err
	}

	subcategory.Inactivate()
	return nil
}

// Suppliers

func (c *Category) GetSuppliers() map[string]string {
	return c.suppliers
}

func (c *Category) AddSupplier(supplierId string) error {
	if !validateStringEmpty(supplierId) {
		return errors.New(ValidateMessageToSupplierId)
	}

	c.suppliers[supplierId] = supplierId
	return nil
}

func (c *Category) RemoveSupplier(supplierId string) error {
	if !validateStringEmpty(supplierId) {
		return errors.New(ValidateMessageToSupplierId)
	}

	delete(c.suppliers, supplierId)
	return nil
}

func NewCategory(name, assistanceType string, id *string) (*Category, error) {
	category := &Category{
		name:           name,
		assistanceType: assistanceType,
		subcategories:  make(map[string]*Subcategory, 0),
		suppliers:      make(map[string]string, 0),
		active:         true,
	}

	if !validateStringEmpty(category.name) {
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

func validateStringEmpty(name string) bool {
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
