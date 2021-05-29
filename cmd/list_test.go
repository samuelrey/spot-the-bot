package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/stretchr/testify/suite"
)

type ListTestSuite struct{ framework.CommandTestSuite }

func (suite *ListTestSuite) SetupTest() {
	suite.CommandTestSuite.SetupTest()
}

// Test that we reply with the expected content given no users have enrolled.
func (suite *ListTestSuite) TestListNoUsers() {
	suite.Replyer.On("Reply", StrListNoUsers).Return(nil)

	List(&suite.Ctx)

	suite.Replyer.AssertCalled(suite.T(), "Reply", StrListNoUsers)
}

// Test that we reply with the expected content given users have enrolled.
func (suite *ListTestSuite) TestListWithUsers() {
	*suite.Ctx.EnrolledUsers = []framework.User{suite.Ctx.User}

	content := fmt.Sprintf(StrListUsersFmt, suite.Ctx.EnrolledUsers)
	suite.Replyer.On("Reply", content).Return(nil)

	List(&suite.Ctx)

	suite.Replyer.AssertCalled(suite.T(), "Reply", content)
}

func TestListCommand(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}
