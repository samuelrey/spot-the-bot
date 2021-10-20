package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ListSuite struct{ CommandSuite }

// Test that we reply with the expected content given no users have enrolled.
func (suite *ListSuite) TestListNoUsers() {
	suite.RotationRepository.On("FindOne", mock.Anything).Return(&suite.Rotation, nil)
	suite.Messager.On("Reply", StrListNoUsers).Return(nil)

	List(&suite.Ctx)

	suite.Messager.AssertCalled(suite.T(), "Reply", StrListNoUsers)
}

// Test that we reply with the expected content given users have enrolled.
func (suite *ListSuite) TestListWithUsers() {
	suite.Rotation.Join(suite.Actor)

	suite.RotationRepository.On("FindOne", mock.Anything).Return(&suite.Rotation, nil)

	content := fmt.Sprintf(StrListUsersFmt, &suite.Rotation)
	suite.Messager.On("Reply", content).Return(nil)

	List(&suite.Ctx)

	suite.Messager.AssertCalled(suite.T(), "Reply", content)
}

func TestListCommand(t *testing.T) {
	suite.Run(t, new(ListSuite))
}
