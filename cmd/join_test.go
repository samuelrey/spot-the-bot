package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JoinTestSuite struct{ framework.CommandTestSuite }

func (suite *JoinTestSuite) SetupTest() {
	suite.Replyer = framework.MockReplyer{}
	enrolledUsers := make([]framework.User, 0)
	suite.User = framework.User{ID: "amethyst#4422", Username: "amethyst"}

	suite.Ctx = framework.Context{
		Replyer:       &suite.Replyer,
		EnrolledUsers: &enrolledUsers,
		User:          suite.User,
	}
}

// Test that the acting user is added to the list of enrolled users.
func (suite *JoinTestSuite) TestJoinUser() {
	content := fmt.Sprintf(StrJoinWelcomeFmt, suite.User)
	suite.Replyer.On("Reply", content).Return(nil)

	Join(&suite.Ctx)

	suite.Replyer.AssertCalled(suite.T(), "Reply", content)
	suite.Require().Equal(
		[]framework.User{suite.User},
		*suite.Ctx.EnrolledUsers,
	)
}

// Test that the acting user is not added again if they are already enrolled.
func (suite *JoinTestSuite) TestJoinUserAlreadyEnrolled() {
	*suite.Ctx.EnrolledUsers = []framework.User{suite.Ctx.User}
	suite.Replyer.On("Reply", mock.Anything).Return(nil)

	Join(&suite.Ctx)

	suite.Replyer.AssertNotCalled(suite.T(), "Reply", mock.Anything)
	suite.Require().Equal(
		[]framework.User{suite.User},
		*suite.Ctx.EnrolledUsers,
	)
}

func TestJoinCommand(t *testing.T) {
	suite.Run(t, new(JoinTestSuite))
}
