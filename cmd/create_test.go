package cmd

import (
	"fmt"
	"testing"

	"github.com/samuelrey/spot-the-bot/framework"
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/stretchr/testify/suite"
)

type CreateTestSuite struct {
	CommandTestSuite
	notActor message.MessageUser
}

func (suite *CreateTestSuite) SetupTest() {
	suite.CommandTestSuite.SetupTest()
	suite.notActor = message.MessageUser{ID: "osh#1219", Username: "osh"}
}

// Test that the acting user can create a playlist and the playlist URL is
// direct messaged to them.
func (suite *CreateTestSuite) TestCreate() {
	suite.UserQueue.Push(suite.Actor)

	playlist := framework.Playlist{
		ID:  suite.Ctx.PlaylistName,
		URL: suite.Ctx.PlaylistName,
	}
	suite.PlaylistCreator.On("CreatePlaylist", suite.Ctx.PlaylistName).Return(playlist, nil)

	content := fmt.Sprintf(StrPlaylistCreatedFmt, suite.Ctx.PlaylistName)
	suite.Messager.On("DirectMessage", suite.Actor.ID, content).Return(nil)

	Create(&suite.Ctx)

	suite.PlaylistCreator.AssertCalled(suite.T(), "CreatePlaylist", suite.Ctx.PlaylistName)
	suite.Messager.AssertCalled(suite.T(), "DirectMessage", suite.Actor.ID, content)
}

// Test that the acting user cannot create a playlist if they are not the head
// of the queue.
func (suite *CreateTestSuite) TestActorNotHeadOfQueue() {
	suite.UserQueue.Push(suite.notActor)
	suite.UserQueue.Push(suite.Actor)
	playlist := framework.Playlist{
		ID:  suite.Ctx.PlaylistName,
		URL: suite.Ctx.PlaylistName,
	}
	suite.PlaylistCreator.On("CreatePlaylist", suite.Ctx.PlaylistName).Return(playlist, nil)

	content := fmt.Sprintf(StrPlaylistCreatedFmt, suite.Ctx.PlaylistName)
	suite.Messager.On("DirectMessage", suite.Actor.ID, content).Return(nil)

	Create(&suite.Ctx)

	suite.PlaylistCreator.AssertNotCalled(suite.T(), "CreatePlaylist", suite.Ctx.PlaylistName)
	suite.Messager.AssertNotCalled(suite.T(), "DirectMessage", suite.Actor.ID, content)
}

// Test that no direct message is sent if the create playlist function returns
// an error.
func (suite *CreateTestSuite) TestNoDirectMessageOnError() {
	suite.UserQueue.Push(suite.Actor)
	suite.Ctx.PlaylistName = "Error"
	suite.PlaylistCreator.On("CreatePlaylist", suite.Ctx.PlaylistName)

	content := fmt.Sprintf(StrPlaylistCreatedFmt, suite.Ctx.PlaylistName)
	suite.Messager.On("DirectMessage", suite.Actor.ID, content).Return(nil)

	Create(&suite.Ctx)

	suite.PlaylistCreator.AssertCalled(suite.T(), "CreatePlaylist", suite.Ctx.PlaylistName)
	suite.Messager.AssertNotCalled(suite.T(), "DirectMessage", suite.Actor.ID, content)
}

func TestCreateCommand(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
