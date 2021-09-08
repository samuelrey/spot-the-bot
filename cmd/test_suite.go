package cmd

import (
	"github.com/samuelrey/spot-the-bot/framework"
	"github.com/stretchr/testify/suite"
)

type CommandTestSuite struct {
	suite.Suite
	Actor           framework.MessageUser
	Ctx             CommandContext
	UserQueue       framework.UserQueue
	Messager        framework.MockMessager
	PlaylistCreator framework.MockPlaylistCreator
}

func (suite *CommandTestSuite) SetupTest() {
	suite.Actor = framework.MessageUser{ID: "amethyst#4422", Username: "amethyst"}
	suite.Messager = framework.MockMessager{}
	suite.PlaylistCreator = framework.MockPlaylistCreator{}

	suite.UserQueue = framework.NewUserQueue([]framework.MessageUser{})

	suite.Ctx = CommandContext{
		Messager:        &suite.Messager,
		PlaylistCreator: &suite.PlaylistCreator,
		PlaylistName:    "Einstok",
		UserQueue:       &suite.UserQueue,
		Actor:           suite.Actor,
	}
}
