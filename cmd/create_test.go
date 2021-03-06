package cmd

import (
	"errors"
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/playlist"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateSuite struct {
	CommandSuite
	notActor message.User
}

func (suite *CreateSuite) SetupTest() {
	suite.CommandSuite.SetupTest()
	suite.notActor = message.User{ID: "osh#1219", Username: "osh"}
}

// Test that the acting user can create a playlist and the playlist URL is
// direct messaged to them.
func (suite *CreateSuite) TestCreate() {
	suite.Rotation.Join(suite.Actor)
	suite.RotationRepository.On("FindOne", mock.Anything).Return(&suite.Rotation, nil)

	playlist := playlist.Playlist{
		ID:  suite.Ctx.PlaylistName,
		URL: suite.Ctx.PlaylistName,
	}
	suite.Creator.On("Create", suite.Ctx.PlaylistName).Return(&playlist, nil)

	content := fmt.Sprintf(StrPlaylistCreatedFmt, suite.Ctx.PlaylistName)
	suite.Messager.On("DirectMessage", suite.Actor.ID, content).Return(nil)

	Create(&suite.Ctx)

	suite.Creator.AssertCalled(suite.T(), "Create", suite.Ctx.PlaylistName)
	suite.Messager.AssertCalled(suite.T(), "DirectMessage", suite.Actor.ID, content)
}

// Test that the acting user cannot create a playlist if they are not the head
// of the queue.
func (suite *CreateSuite) TestActorNotHeadOfQueue() {
	suite.Rotation.Join(suite.notActor)
	suite.Rotation.Join(suite.Actor)
	playlist := playlist.Playlist{
		ID:  suite.Ctx.PlaylistName,
		URL: suite.Ctx.PlaylistName,
	}
	suite.RotationRepository.On("FindOne", mock.Anything).Return(&suite.Rotation, nil)
	suite.Creator.On("Create", suite.Ctx.PlaylistName).Return(playlist, nil)

	content := fmt.Sprintf(StrPlaylistCreatedFmt, suite.Ctx.PlaylistName)
	suite.Messager.On("DirectMessage", suite.Actor.ID, content).Return(nil)

	Create(&suite.Ctx)

	suite.Creator.AssertNotCalled(suite.T(), "Create", suite.Ctx.PlaylistName)
	suite.Messager.AssertNotCalled(suite.T(), "DirectMessage", suite.Actor.ID, content)
}

// Test that no direct message is sent if the create playlist function returns
// an error.
func (suite *CreateSuite) TestNoDirectMessageOnError() {
	suite.Rotation.Join(suite.Actor)
	suite.Ctx.PlaylistName = "Error"
	suite.RotationRepository.On("FindOne", mock.Anything).Return(&suite.Rotation, nil)
	suite.Creator.On("Create", suite.Ctx.PlaylistName).Return(nil, errors.New(""))

	content := fmt.Sprintf(StrPlaylistCreatedFmt, suite.Ctx.PlaylistName)
	suite.Messager.On("DirectMessage", suite.Actor.ID, content).Return(nil)

	Create(&suite.Ctx)

	suite.Creator.AssertCalled(suite.T(), "Create", suite.Ctx.PlaylistName)
	suite.Messager.AssertNotCalled(suite.T(), "DirectMessage", suite.Actor.ID, content)
}

func TestCreateCommand(t *testing.T) {
	suite.Run(t, new(CreateSuite))
}
