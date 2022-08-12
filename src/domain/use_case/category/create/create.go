package create

import (
	"errors"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	infra "github.com/WilkerAlves/assistance-go/src/domain/interface/service"
	"github.com/WilkerAlves/assistance-go/src/domain/service"
)

type createCategoryUseCase struct {
	eventService    infra.IEventService
	categoryService service.ICategoryService
	generateIds     infra.IGeneratedIds
}

func (c *createCategoryUseCase) Execute(input InputCrateCategory) error {

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

func NewCreateCategoryUseCase(es infra.IEventService, cs service.ICategoryService, gIds infra.IGeneratedIds) *createCategoryUseCase {
	return &createCategoryUseCase{
		eventService:    es,
		categoryService: cs,
		generateIds:     gIds,
	}
}
