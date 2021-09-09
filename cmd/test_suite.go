package cmd

import (
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/playlist"
	"github.com/samuelrey/spot-the-bot/rotation"
	"github.com/stretchr/testify/suite"
)

type CommandTestSuite struct {
	suite.Suite
	Actor           message.User
	Ctx             CommandContext
	UserQueue       rotation.Rotation
	Messager        message.MockMessager
	PlaylistCreator playlist.MockPlaylistCreator
}

func (suite *CommandTestSuite) SetupTest() {
	suite.Actor = message.User{ID: "amethyst#4422", Username: "amethyst"}
	suite.Messager = message.MockMessager{}
	suite.PlaylistCreator = playlist.MockPlaylistCreator{}

	suite.UserQueue = rotation.NewRotation([]message.User{})

	suite.Ctx = CommandContext{
		Messager:        &suite.Messager,
		PlaylistCreator: &suite.PlaylistCreator,
		PlaylistName:    "Einstok",
		UserQueue:       &suite.UserQueue,
		Actor:           suite.Actor,
	}
}
