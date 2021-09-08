package playlist

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

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
