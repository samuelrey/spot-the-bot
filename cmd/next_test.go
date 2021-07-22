package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/framework"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type NextTestSuite struct {
	framework.CommandTestSuite
	notActor framework.MessageUser
}

func (suite *NextTestSuite) SetupTest() {
	suite.CommandTestSuite.SetupTest()
	suite.notActor = framework.MessageUser{ID: "osh#1219", Username: "osh"}
}

// Test that we do not pop/push the user at the head if it is not the actor.
func (suite *NextTestSuite) TestActorNotHeadOfQueue() {
	suite.UserQueue.Push(suite.notActor)
	suite.UserQueue.Push(suite.Actor)
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Next(&suite.Ctx)

	suite.Messager.AssertNotCalled(suite.T(), "Reply", mock.Anything)
	expected := framework.NewUserQueue([]framework.MessageUser{suite.notActor, suite.Actor})
	suite.Require().Equal(expected, suite.UserQueue)
}

// Test that we pop/push the user at the head if it is the actor.
func (suite *NextTestSuite) TestNext() {
	suite.UserQueue.Push(suite.Actor)
	suite.UserQueue.Push(suite.notActor)
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Next(&suite.Ctx)

	content := fmt.Sprintf(StrSkipUser, suite.Actor)
	suite.Messager.AssertCalled(suite.T(), "Reply", content)
	content = fmt.Sprintf(StrNextUser, suite.notActor)
	suite.Messager.AssertCalled(suite.T(), "Reply", content)
	expected := framework.NewUserQueue([]framework.MessageUser{suite.notActor, suite.Actor})
	suite.Require().Equal(expected, suite.UserQueue)
}

func TestNextCommand(t *testing.T) {
	suite.Run(t, new(NextTestSuite))
}
