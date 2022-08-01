package service

type IEventService interface {
	Send(event string, body interface{}) error
	//Read(event string, body interface{}) entity.Category
}

//
//type EventCategory struct {
//	category entity.Category
//}
//
//func (e *EventCategory) Send() error  {
//
//}
//
//func (e *EventCategory) Read() entity.Category  {
//
//}
