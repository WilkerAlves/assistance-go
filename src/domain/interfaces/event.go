package interfaces

type IEventService interface {
	Send(event string, body interface{}) bool
}
