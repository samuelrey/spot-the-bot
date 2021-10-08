package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/rotation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JoinSuite struct{ CommandSuite }

// Test that the acting user is added to the list of enrolled users.
func (suite *JoinSuite) TestJoinUser() {
	content := fmt.Sprintf(StrJoinFmt, suite.Actor)
	suite.Messager.On("Reply", content).Return(nil)

	Join(&suite.Ctx)

	suite.Messager.AssertCalled(suite.T(), "Reply", content)

	expected := rotation.NewRotation([]message.User{suite.Actor})
	suite.Require().Equal(expected, suite.UserQueue)
}

// Test that the acting user is not added again if they are already enrolled.
func (suite *JoinSuite) TestJoinUserAlreadyEnrolled() {
	suite.UserQueue.Next(suite.Actor)
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Join(&suite.Ctx)

	suite.Messager.AssertNotCalled(suite.T(), "Reply", mock.Anything)

	expected := rotation.NewRotation([]message.User{suite.Actor})
	suite.Require().Equal(expected, suite.UserQueue)
}

func TestJoinCommand(t *testing.T) {
	suite.Run(t, new(JoinSuite))
}
