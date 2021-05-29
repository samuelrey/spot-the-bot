package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LeaveTestSuite struct {
	framework.CommandTestSuite
	notActor framework.User
}

func (suite *LeaveTestSuite) SetupTest() {
	suite.CommandTestSuite.SetupTest()
	suite.notActor = framework.User{ID: "osh#1219", Username: "osh"}
	suite.EnrolledUsers = []framework.User{suite.notActor}
}

// Test that we do not remove any users if the actor is not enrolled.
func (suite *LeaveTestSuite) TestLeaveUserNotEnrolled() {
	suite.Replyer.On("Reply", mock.Anything).Return(nil)

	Leave(&suite.Ctx)

	suite.Replyer.AssertNotCalled(suite.T(), "Reply", mock.Anything)
	suite.Require().Equal(
		[]framework.User{suite.notActor},
		suite.EnrolledUsers,
	)
}

// Test that we only remove the actor if they are enrolled.
func (suite *LeaveTestSuite) TestLeaveUser() {
	suite.EnrolledUsers = append(suite.EnrolledUsers, suite.Actor)

	content := fmt.Sprintf(StrLeaveFmt, suite.Actor)
	suite.Replyer.On("Reply", content).Return(nil)

	Leave(&suite.Ctx)

	suite.Replyer.AssertCalled(suite.T(), "Reply", content)
	suite.Require().Equal(
		[]framework.User{suite.notActor},
		suite.EnrolledUsers,
	)
}

func TestLeaveCommand(t *testing.T) {
	suite.Run(t, new(LeaveTestSuite))
}
