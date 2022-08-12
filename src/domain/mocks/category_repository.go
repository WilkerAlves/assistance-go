package mocks

import (
	"errors"
	"fmt"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MyMockedCategoryRepository struct {
	mock.Mock
	DB []entity.Category
}

func (m *MyMockedCategoryRepository) Create(category entity.Category) error {
	m.DB = append(m.DB, category)
	return nil
}

func (m *MyMockedCategoryRepository) Update(category entity.Category) error {
	for i, oldCategory := range m.DB {
		if oldCategory.GetID() == category.GetID() {
			err := oldCategory.ChangeName(category.GetName())
			if err != nil {
				return err
			}

			err = oldCategory.ChangeAssistanceType(category.GetAssistanceType())
			if err != nil {
				return err
			}

			if category.GetStatus() == false {
				oldCategory.Inactivate()
			}

			m.DB[i] = oldCategory

			return nil
		}
	}
	return errors.New(fmt.Sprintf("category not found. id: %s", category.GetID()))
}

func (m *MyMockedCategoryRepository) Find(id string) (*entity.Category, error) {
	for _, category := range m.DB {
		if category.GetID() == id {
			return &category, nil
		}
	}
	return nil, nil
}

func (m *MyMockedCategoryRepository) FindByName(name string) (*entity.Category, error) {
	for _, category := range m.DB {
		if category.GetName() == name {
			return &category, nil
		}
	}
	return nil, nil
}

func (m *MyMockedCategoryRepository) FindAll(active *bool) ([]*entity.Category, error) {
	output := make([]*entity.Category, 0)
	if active == nil {
		for _, category := range m.DB {
			output = append(output, &category)
		}

		return output, nil
	}

	for i := range m.DB {
		if m.DB[i].GetStatus() == *active {
			output = append(output, &m.DB[i])
		}
	}

	return output, nil
}
