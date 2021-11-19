package cmd

import (
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/mocks"
	"github.com/stretchr/testify/suite"
)

type CommandSuite struct {
	suite.Suite
	Actor              message.User
	Ctx                Context
	Rotation           message.Rotation
	Messager           mocks.Messager
	Creator            mocks.Creator
	RepositoryProvider mocks.IProvider
	RotationRepository mocks.IRotationRepository
}

func (suite *CommandSuite) SetupTest() {
	suite.Actor = message.User{ID: "amethyst#4422", Username: "amethyst"}
	suite.Messager = mocks.Messager{}
	suite.Creator = mocks.Creator{}
	suite.RotationRepository = mocks.IRotationRepository{}
	suite.RepositoryProvider = mocks.IProvider{}
	suite.RepositoryProvider.On("GetRotationRepository").Return(&suite.RotationRepository)
	suite.Rotation = message.NewRotation([]message.User{}, "einstok")

	suite.Ctx = Context{
		Messager:           &suite.Messager,
		PlaylistCreator:    &suite.Creator,
		PlaylistName:       "Einstok",
		Actor:              suite.Actor,
		RepositoryProvider: &suite.RepositoryProvider,
		ServerID:           "Einstok",
	}
}
