package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/message"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JoinSuite struct{ CommandSuite }

// Test that the acting user is added to the list of enrolled users.
func (suite *JoinSuite) TestJoinUser() {
	suite.RotationRepository.On("FindOne", mock.Anything).Return(&suite.Rotation, nil)
	suite.RotationRepository.On("Upsert", mock.Anything).Return(nil)

	content := fmt.Sprintf(StrJoinFmt, suite.Actor)
	suite.Messager.On("Reply", content).Return(nil)

	Join(&suite.Ctx)

	suite.Messager.AssertCalled(suite.T(), "Reply", content)

	expected := message.NewRotation([]message.User{suite.Actor}, "einstok")
	suite.Require().Equal(expected, suite.Rotation)
}

// Test that the acting user is not added again if they are already enrolled.
func (suite *JoinSuite) TestJoinUserAlreadyEnrolled() {
	suite.Rotation.Join(suite.Actor)
	suite.RotationRepository.On("FindOne", mock.Anything).Return(&suite.Rotation, nil)
	suite.RotationRepository.On("Upsert", mock.Anything).Return(nil)
	suite.Messager.On("Reply", mock.Anything).Return(nil)

	Join(&suite.Ctx)

	suite.Messager.AssertNotCalled(suite.T(), "Reply", mock.Anything)

	expected := message.NewRotation([]message.User{suite.Actor}, "einstok")
	suite.Require().Equal(expected, suite.Rotation)
}

func TestJoinCommand(t *testing.T) {
	suite.Run(t, new(JoinSuite))
}
