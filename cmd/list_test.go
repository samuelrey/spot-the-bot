package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/stretchr/testify/suite"
)

type ListTestSuite struct{ framework.CommandTestSuite }

// Test that we reply with the expected content given no users have enrolled.
func (suite *ListTestSuite) TestListNoUsers() {
	suite.Replyer.On("Reply", StrListNoUsers).Return(nil)

	List(&suite.Ctx)

	suite.Replyer.AssertCalled(suite.T(), "Reply", StrListNoUsers)
}

// Test that we reply with the expected content given users have enrolled.
func (suite *ListTestSuite) TestListWithUsers() {
	suite.EnrolledUsers = []framework.User{suite.Actor}

	content := fmt.Sprintf(StrListUsersFmt, suite.EnrolledUsers)
	suite.Replyer.On("Reply", content).Return(nil)

	List(&suite.Ctx)

	suite.Replyer.AssertCalled(suite.T(), "Reply", content)
}

func TestListCommand(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}
