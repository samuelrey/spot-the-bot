package framework

import (
	"errors"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockMessager struct{ mock.Mock }

func (mm *MockMessager) Reply(content string) error {
	mm.Called(content)
	return nil
}

func (mm *MockMessager) DirectMessage(recipient, content string) error {
	mm.Called(recipient, content)
	return nil
}

type MockPlaylistCreator struct{ mock.Mock }

func (mp *MockPlaylistCreator) CreatePlaylist(playlistName string) (*Playlist, error) {
	mp.Called(playlistName)
	if playlistName == "Error" {
		return nil, errors.New("Error")
	}

	return &Playlist{
		ID:  playlistName,
		URL: playlistName,
	}, nil
}

type CommandTestSuite struct {
	suite.Suite
	Actor           MessageUser
	Ctx             CommandContext
	UserQueue       UserQueue
	Messager        MockMessager
	PlaylistCreator MockPlaylistCreator
}

func (suite *CommandTestSuite) SetupTest() {
	suite.Actor = MessageUser{ID: "amethyst#4422", Username: "amethyst"}
	suite.Messager = MockMessager{}
	suite.PlaylistCreator = MockPlaylistCreator{}

	q := NewSimpleUserQueue([]MessageUser{})
	suite.UserQueue = &q

	suite.Ctx = CommandContext{
		Messager:        &suite.Messager,
		PlaylistCreator: &suite.PlaylistCreator,
		PlaylistName:    "Einstok",
		UserQueue:       suite.UserQueue,
		Actor:           suite.Actor,
	}
}
