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
	Actor         User
	Ctx           Context
	EnrolledUsers []User
	Replyer       MockReplyer
}

func (suite *CommandTestSuite) SetupTest() {
	suite.Actor = User{ID: "amethyst#4422", Username: "amethyst"}
	suite.Replyer = MockReplyer{}
	suite.EnrolledUsers = make([]User, 0)

	suite.Ctx = Context{
		Replyer:       &suite.Replyer,
		EnrolledUsers: &suite.EnrolledUsers,
		Actor:         suite.Actor,
	}
}
