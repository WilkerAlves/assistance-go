package entity

type Subcategory struct {
	category   *Category
	active     bool
	stockGroup string
}

func (s *Subcategory) GetStatus() bool {
	return s.active
}

func (s *Subcategory) Inactivate() {
	s.active = false
}

func (s *Subcategory) ChangeStockGroup(groupStock string) {
	s.stockGroup = groupStock
}

func (s *Subcategory) GetStockGroup() string {
	return s.stockGroup
}

func (s *Subcategory) GetName() string {
	return s.category.name
}
