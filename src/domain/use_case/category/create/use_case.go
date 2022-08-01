package create

import (
	"github.com/WilkerAlves/assistance-go/src/domain/entity"
	infra "github.com/WilkerAlves/assistance-go/src/domain/interface/service"
	"github.com/WilkerAlves/assistance-go/src/domain/service"
)

func Execute(
	input InputCrateCategory,
	eventService infra.IEventService,
	categoryService service.ICategoryService,
) error {

	category, err := entity.NewCategory(input.Name, input.AssistanceType, nil)
	if err != nil {
		return err
	}

	err = categoryService.Create(*category)
	if err != nil {
		return err
	}

	err = eventService.Send("CREATE_CATEGORY", category)
	if err != nil {
		return err
	}

	return nil
}
