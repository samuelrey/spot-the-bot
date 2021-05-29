package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ListTestSuite struct {
	suite.Suite
	ctx     framework.Context
	replyer MockReplyer
}

type MockReplyer struct{ mock.Mock }

func (m *MockReplyer) Reply(content string) error {
	m.Called(content)
	return nil
}

func (suite *ListTestSuite) SetupTest() {
	suite.replyer = MockReplyer{}
	enrolledUsers := make([]framework.User, 0)
	user := framework.User{ID: "amethyst#4422", Username: "amethyst"}

	suite.ctx = framework.Context{
		Replyer:       &suite.replyer,
		EnrolledUsers: &enrolledUsers,
		User:          user,
	}
}

// Test that we reply with the expected content given no users have enrolled.
func (suite *ListTestSuite) TestListNoUsers() {
	suite.replyer.On("Reply", StrListNoUsers).Return(nil)

	List(&suite.ctx)

	suite.replyer.AssertCalled(suite.T(), "Reply", StrListNoUsers)
}

// Test that we reply with the expected content given users have enrolled.
func (suite *ListTestSuite) TestListWithUsers() {
	*suite.ctx.EnrolledUsers = []framework.User{suite.ctx.User}

	content := fmt.Sprintf(StrListUsersFmt, suite.ctx.EnrolledUsers)
	suite.replyer.On("Reply", content).Return(nil)

	List(&suite.ctx)

	suite.replyer.AssertCalled(suite.T(), "Reply", content)
}

func TestListCommand(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}
