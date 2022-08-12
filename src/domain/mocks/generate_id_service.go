package mocks

import "github.com/stretchr/testify/mock"

type MyMockedGeneratedIdsService struct {
	mock.Mock
}

func (m *MyMockedGeneratedIdsService) Create() (string, error) {
	return "12345677", nil
}
