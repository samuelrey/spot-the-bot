package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JoinTestSuite struct{ framework.CommandTestSuite }

// Test that the acting user is added to the list of enrolled users.
func (suite *JoinTestSuite) TestJoinUser() {
	content := fmt.Sprintf(StrJoinFmt, suite.Actor)
	suite.Replyer.On("Reply", content).Return(nil)

	Join(&suite.Ctx)

	suite.Replyer.AssertCalled(suite.T(), "Reply", content)
	suite.Require().Equal(
		[]framework.MessageUser{suite.Actor},
		suite.EnrolledUsers,
	)
}

// Test that the acting user is not added again if they are already enrolled.
func (suite *JoinTestSuite) TestJoinUserAlreadyEnrolled() {
	suite.EnrolledUsers = []framework.MessageUser{suite.Actor}
	suite.Replyer.On("Reply", mock.Anything).Return(nil)

	Join(&suite.Ctx)

	suite.Replyer.AssertNotCalled(suite.T(), "Reply", mock.Anything)
	suite.Require().Equal(
		[]framework.MessageUser{suite.Actor},
		suite.EnrolledUsers,
	)
}

func TestJoinCommand(t *testing.T) {
	suite.Run(t, new(JoinTestSuite))
}
