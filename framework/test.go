package framework

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockMessager struct{ mock.Mock }

func (m *MockMessager) Reply(content string) error {
	m.Called(content)
	return nil
}

func (m *MockMessager) DirectMessage(recipient, content string) error {
	m.Called(recipient, content)
	return nil
}

type CommandTestSuite struct {
	suite.Suite
	Actor         MessageUser
	Ctx           CommandContext
	EnrolledUsers []MessageUser
	Replyer       MockMessager
}

func (suite *CommandTestSuite) SetupTest() {
	suite.Actor = MessageUser{ID: "amethyst#4422", Username: "amethyst"}
	suite.Replyer = MockMessager{}
	suite.EnrolledUsers = make([]MessageUser, 0)

	suite.Ctx = CommandContext{
		Messager:      &suite.Replyer,
		EnrolledUsers: &suite.EnrolledUsers,
		Actor:         suite.Actor,
	}
}
