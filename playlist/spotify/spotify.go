package spotify

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/samuelrey/spot-the-bot/playlist"
	"github.com/zmb3/spotify"
)

type PlaylistCreator struct {
	*spotify.Client
	user *spotify.User
}

// NewPlaylistCreator takes a user through the Spotify authorization flow and
// returns a PlaylistCreator which is tied to a single Spotify user's account.
// We recommend creating an account on Spotify specific for this bot.
func NewPlaylistCreator(conf SpotifyConfig) (playlist.PlaylistCreator, error) {
	log.Println("Auth server starting.")
	a := newAuthenticator(conf)
	srv := a.startAuthServer()
	defer func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println(err)
		} else {
			log.Println("Auth server shutdown.")
		}
	}()

	log.Printf("Navigate here to authorize Spotify user: %s\n", a.authURL)
	token, err := getToken()
	if err != nil {
		return nil, errors.Wrap(err, "authorize spotify user")
	}

	client := a.NewClient(token)
	spotifyUser, err := client.CurrentUser()
	if err != nil {
		return nil, errors.Wrap(err, "authorize spotify user")
	}

	return &PlaylistCreator{
		Client: &client,
		user:   &spotifyUser.User,
	}, nil
}

// CreatePlaylist creates a playlist an authorized user.
func (sp *PlaylistCreator) CreatePlaylist(playlistName string) (*playlist.Playlist, error) {
	p, err := sp.CreateCollaborativePlaylistForUser(
		sp.user.ID, playlistName, "")
	if err != nil {
		return nil, errors.Wrap(err, "Create Spotify playlist")
	}

	return &playlist.Playlist{
		ID:  p.ID.String(),
		URL: p.ExternalURLs["spotify"],
	}, nil
}