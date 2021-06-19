package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-discord/framework"
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
func (suite *NextTestSuite) TestNextActorNotHeadOfQueue() {
	suite.EnrolledUsers = []framework.MessageUser{suite.notActor, suite.Actor}
	suite.Replyer.On("Reply", mock.Anything).Return(nil)

	Next(&suite.Ctx)

	suite.Replyer.AssertNotCalled(suite.T(), "Reply", mock.Anything)
	suite.Require().Equal(
		[]framework.MessageUser{suite.notActor, suite.Actor},
		suite.EnrolledUsers,
	)
}

// Test that we pop/push the user at the head if it is the actor.
func (suite *NextTestSuite) TestNextActor() {
	suite.EnrolledUsers = []framework.MessageUser{suite.Actor, suite.notActor}
	suite.Replyer.On("Reply", mock.Anything).Return(nil)

	Next(&suite.Ctx)

	content := fmt.Sprintf(StrSkipUser, suite.Actor)
	suite.Replyer.AssertCalled(suite.T(), "Reply", content)
	content = fmt.Sprintf(StrNextUser, suite.notActor)
	suite.Replyer.AssertCalled(suite.T(), "Reply", content)
	suite.Require().Equal(
		[]framework.MessageUser{suite.notActor, suite.Actor},
		suite.EnrolledUsers,
	)
}

func TestNextCommand(t *testing.T) {
	suite.Run(t, new(NextTestSuite))
}
