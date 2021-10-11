package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/rotation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LeaveSuite struct {
	CommandSuite
	notActor message.User
}

func (suite *LeaveSuite) SetupTest() {
	suite.CommandSuite.SetupTest()
	suite.notActor = message.User{ID: "osh#1219", Username: "osh"}
	suite.UserQueue.Join(suite.notActor)
}

// Test that we do not remove any users if the actor is not enrolled.
func (suite *LeaveSuite) TestLeaveUserNotEnrolled() {
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Leave(&suite.Ctx)

	suite.Messager.AssertNotCalled(suite.T(), "Reply", mock.Anything)
	expected := rotation.NewRotation([]message.User{suite.notActor})
	suite.Require().Equal(expected, suite.UserQueue)
}

// Test that we only remove the actor if they are enrolled.
func (suite *LeaveSuite) TestLeaveUser() {
	suite.UserQueue.Join(suite.Actor)

	content := fmt.Sprintf(StrLeaveFmt, suite.Actor)
	suite.Messager.On("Reply", content).Return(nil)

	Leave(&suite.Ctx)

	suite.Messager.AssertCalled(suite.T(), "Reply", content)
	expected := rotation.NewRotation([]message.User{suite.notActor})
	suite.Require().Equal(expected, suite.UserQueue)
}

func TestLeaveCommand(t *testing.T) {
	suite.Run(t, new(LeaveSuite))
}
