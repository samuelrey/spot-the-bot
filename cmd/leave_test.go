package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/framework"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LeaveTestSuite struct {
	framework.CommandTestSuite
	notActor framework.MessageUser
}

func (suite *LeaveTestSuite) SetupTest() {
	suite.CommandTestSuite.SetupTest()
	suite.notActor = framework.MessageUser{ID: "osh#1219", Username: "osh"}
	suite.UserQueue.Push(suite.notActor)
}

// Test that we do not remove any users if the actor is not enrolled.
func (suite *LeaveTestSuite) TestLeaveUserNotEnrolled() {
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Leave(&suite.Ctx)

	suite.Messager.AssertNotCalled(suite.T(), "Reply", mock.Anything)
	expected := framework.NewUserQueue([]framework.MessageUser{suite.notActor})
	suite.Require().Equal(expected, suite.UserQueue)
}

// Test that we only remove the actor if they are enrolled.
func (suite *LeaveTestSuite) TestLeaveUser() {
	suite.UserQueue.Push(suite.Actor)

	content := fmt.Sprintf(StrLeaveFmt, suite.Actor)
	suite.Messager.On("Reply", content).Return(nil)

	Leave(&suite.Ctx)

	suite.Messager.AssertCalled(suite.T(), "Reply", content)
	expected := framework.NewUserQueue([]framework.MessageUser{suite.notActor})
	suite.Require().Equal(expected, suite.UserQueue)
}

func TestLeaveCommand(t *testing.T) {
	suite.Run(t, new(LeaveTestSuite))
}
