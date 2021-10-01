package message

import "github.com/stretchr/testify/mock"

type MockMessager struct{ mock.Mock }

func (mm *MockMessager) Reply(content string) error {
	mm.Called(content)
	return nil
}

func (mm *MockMessager) DirectMessage(recipient, content string) error {
	mm.Called(recipient, content)
	return nil
}
