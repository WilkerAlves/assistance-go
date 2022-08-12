package create

import (
	"errors"

	"github.com/WilkerAlves/assistance-go/src/domain/dto"
	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	"github.com/WilkerAlves/assistance-go/src/domain/interfaces"
)

type createCategoryUseCase struct {
	eventService    interfaces.IEventService
	categoryService interfaces.ICategoryService
	generateIds     interfaces.IGeneratedIds
}

func (c *createCategoryUseCase) Execute(input dto.InputCrateCategory) error {

	id, err := c.generateIds.Create()
	if err != nil {
		return err
	}

	category, err := entity.NewCategory(input.Name, input.AssistanceType, "1234", &id)
	if err != nil {
		return err
	}

	err = c.categoryService.Create(*category)
	if err != nil {
		return err
	}

	if !c.eventService.Send("CREATE_CATEGORY", category) {
		return errors.New("error while dispatch event")
	}

	return nil
}

func NewCreateCategoryUseCase(es interfaces.IEventService, cs interfaces.ICategoryService, gIds interfaces.IGeneratedIds) *createCategoryUseCase {
	return &createCategoryUseCase{
		eventService:    es,
		categoryService: cs,
		generateIds:     gIds,
	}
}
