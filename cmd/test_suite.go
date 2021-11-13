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
	RotationRepository mocks.RotationUpserterFinder
}

func (suite *CommandSuite) SetupTest() {
	suite.Actor = message.User{ID: "amethyst#4422", Username: "amethyst"}
	suite.Messager = mocks.Messager{}
	suite.Creator = mocks.Creator{}
	suite.RotationRepository = mocks.RotationUpserterFinder{}
	suite.Rotation = message.NewRotation([]message.User{}, "einstok")

	suite.Ctx = Context{
		Messager:           &suite.Messager,
		PlaylistCreator:    &suite.Creator,
		PlaylistName:       "Einstok",
		Actor:              suite.Actor,
		RotationRepository: &suite.RotationRepository,
		ServerID:           "Einstok",
	}
}
