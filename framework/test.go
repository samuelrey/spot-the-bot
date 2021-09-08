package framework

import (
	"errors"

	"github.com/stretchr/testify/mock"
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
