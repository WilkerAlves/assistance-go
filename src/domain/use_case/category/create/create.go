package create

import (
	"errors"

	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	infra "github.com/WilkerAlves/assistance-go/src/domain/interface/service"
	"github.com/WilkerAlves/assistance-go/src/domain/service"
)

type CreateCategoryUseCase struct {
	EventService    infra.IEventService
	CategoryService service.ICategoryService
	GenerateIds     infra.IGeneratedIds
}

func (c *CreateCategoryUseCase) Execute(input InputCrateCategory) error {

	id, err := c.GenerateIds.Create()
	if err != nil {
		return err
	}

	category, err := entity.NewCategory(input.Name, input.AssistanceType, "1234", &id)
	if err != nil {
		return err
	}

	err = c.CategoryService.Create(*category)
	if err != nil {
		return err
	}

	if !c.EventService.Send("CREATE_CATEGORY", category) {
		return errors.New("error while dispatch event")
	}

	return nil
}
