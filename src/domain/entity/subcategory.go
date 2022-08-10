package entity

type Subcategory struct {
	category *Category
	active   bool
}

func (s *Subcategory) GetStatus() bool {
	return s.active
}

func (s *Subcategory) Inactivate() {
	s.active = false
}
