package mocks

import "github.com/stretchr/testify/mock"

type MyMockedEventService struct {
	mock.Mock
}

func (m *MyMockedEventService) Send(eventName string, body interface{}) bool {
	return true
}
