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
	Actor         User
	Ctx           Context
	EnrolledUsers []User
	Replyer       MockMessager
}

func (suite *CommandTestSuite) SetupTest() {
	suite.Actor = User{ID: "amethyst#4422", Username: "amethyst"}
	suite.Replyer = MockMessager{}
	suite.EnrolledUsers = make([]User, 0)

	suite.Ctx = Context{
		Messager:      &suite.Replyer,
		EnrolledUsers: &suite.EnrolledUsers,
		Actor:         suite.Actor,
	}
}
