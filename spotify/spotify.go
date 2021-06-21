package spotify

import (
	"github.com/pkg/errors"
	"github.com/samuelrey/spot-discord/framework"
	"github.com/zmb3/spotify"
)

type SpotifyPlaylistCreator struct {
	client *spotify.Client
	user   *spotify.User
}

// CreatePlaylist creates a playlist an authorized user.
func (sp *SpotifyPlaylistCreator) CreatePlaylist(playlistName string) (*framework.Playlist, error) {
	playlist, err := sp.client.CreateCollaborativePlaylistForUser(sp.user.ID, playlistName, "")
	if err != nil {
		return nil, errors.Wrap(err, "Create Spotify playlist")
	}

	return &framework.Playlist{
		ID:  playlist.ID.String(),
		URL: playlist.ExternalURLs["spotify"],
	}, nil
}
