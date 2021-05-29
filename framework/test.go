package framework

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockReplyer struct{ mock.Mock }

func (m *MockReplyer) Reply(content string) error {
	m.Called(content)
	return nil
}

type CommandTestSuite struct {
	suite.Suite
	Ctx     Context
	Replyer MockReplyer
	User    User
}
