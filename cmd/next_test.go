package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/rotation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type NextSuite struct {
	CommandSuite
	notActor message.User
}

func (suite *NextSuite) SetupTest() {
	suite.CommandSuite.SetupTest()
	suite.notActor = message.User{ID: "osh#1219", Username: "osh"}
}

// Test that we do not pop/push the user at the head if it is not the actor.
func (suite *NextSuite) TestActorNotHeadOfQueue() {
	suite.UserQueue.Push(suite.notActor)
	suite.UserQueue.Push(suite.Actor)
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Next(&suite.Ctx)

	suite.Messager.AssertNotCalled(suite.T(), "Reply", mock.Anything)
	expected := rotation.NewRotation([]message.User{suite.notActor, suite.Actor})
	suite.Require().Equal(expected, suite.UserQueue)
}

// Test that we pop/push the user at the head if it is the actor.
func (suite *NextSuite) TestNext() {
	suite.UserQueue.Push(suite.Actor)
	suite.UserQueue.Push(suite.notActor)
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Next(&suite.Ctx)

	content := fmt.Sprintf(StrSkipUser, suite.Actor)
	suite.Messager.AssertCalled(suite.T(), "Reply", content)
	content = fmt.Sprintf(StrNextUser, suite.notActor)
	suite.Messager.AssertCalled(suite.T(), "Reply", content)
	expected := rotation.NewRotation([]message.User{suite.notActor, suite.Actor})
	suite.Require().Equal(expected, suite.UserQueue)
}

func TestNextCommand(t *testing.T) {
	suite.Run(t, new(NextSuite))
}
