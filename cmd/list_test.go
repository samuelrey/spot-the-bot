package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ListSuite struct{ CommandSuite }

// Test that we reply with the expected content given no users have enrolled.
func (suite *ListSuite) TestListNoUsers() {
	suite.Messager.On("Reply", StrListNoUsers).Return(nil)

	List(&suite.Ctx)

	suite.Messager.AssertCalled(suite.T(), "Reply", StrListNoUsers)
}

// Test that we reply with the expected content given users have enrolled.
func (suite *ListSuite) TestListWithUsers() {
	suite.UserQueue.Push(suite.Actor)

	content := fmt.Sprintf(StrListUsersFmt, suite.UserQueue)
	suite.Messager.On("Reply", content).Return(nil)

	List(&suite.Ctx)

	suite.Messager.AssertCalled(suite.T(), "Reply", content)
}

func TestListCommand(t *testing.T) {
	suite.Run(t, new(ListSuite))
}
