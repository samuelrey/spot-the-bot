package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/message"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RotateSuite struct {
	CommandSuite
	notActor message.User
}

func (suite *RotateSuite) SetupTest() {
	suite.CommandSuite.SetupTest()
	suite.notActor = message.User{ID: "osh#1219", Username: "osh"}
}

// Test that we pop/Next the user at the head if it is the actor.
func (suite *RotateSuite) TestRotate() {
	suite.Rotation.Join(suite.Actor)
	suite.Rotation.Join(suite.notActor)
	suite.RotationRepository.On("FindOne", mock.Anything).Return(&suite.Rotation, nil)
	suite.RotationRepository.On("Upsert", mock.Anything).Return(nil)
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Rotate(&suite.Ctx)

	content := fmt.Sprintf(StrNextUser, suite.Actor, &suite.notActor)
	suite.Messager.AssertCalled(suite.T(), "Reply", content)
	expected := message.NewRotation([]message.User{suite.notActor, suite.Actor}, "einstok")
	suite.Require().Equal(expected, suite.Rotation)
}

func TestNextCommand(t *testing.T) {
	suite.Run(t, new(RotateSuite))
}
