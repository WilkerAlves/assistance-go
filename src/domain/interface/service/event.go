package service

type IEventService interface {
	Send(event string, body interface{}) bool
}
