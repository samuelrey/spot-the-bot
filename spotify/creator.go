package spotify

import (
	"github.com/pkg/errors"
	"github.com/samuelrey/spot-the-bot/playlist"
	"github.com/zmb3/spotify"
)

type Creator struct {
	*spotify.Client
	user *spotify.User
}

// NewCreator takes a user through the Spotify authorization flow and
// returns a PlaylistCreator which is tied to a single Spotify user's account.
// We recommend creating an account on Spotify specific for this bot.
func NewCreator(conf SpotifyConfig) (playlist.Creator, error) {
	client, err := newSpotifyClient(conf)	
	if err != nil {
		return nil, errors.Wrap(err, "cannot create client")
	}

	spotifyUser, err := client.CurrentUser()
	if err != nil {
		return nil, errors.Wrap(err, "authorize spotify user")
	}

	return &Creator{
		Client: client,
		user:   &spotifyUser.User,
	}, nil
}

// CreatePlaylist creates a playlist an authorized user.
func (sp *Creator) Create(name string) (*playlist.Playlist, error) {
	p, err := sp.CreateCollaborativePlaylistForUser(sp.user.ID, name, "")
	if err != nil {
		return nil, errors.Wrap(err, "Create Spotify playlist")
	}

	return &playlist.Playlist{
		ID:  p.ID.String(),
		URL: p.ExternalURLs["spotify"],
	}, nil
}
