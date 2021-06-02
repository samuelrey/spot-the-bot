package spotify

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/zmb3/spotify"
)

var (
	authURL              string
	config               *Config
	spotifyAuthenticator spotify.Authenticator
	tknHandler           *TokenHandler
)

const (
	StrAuthMessageFmt = "Click this link so we can create your playlist on Spotify:\n%s"
)

func init() {
	config = LoadConfig("secrets_spotify.json")
	spotifyAuthenticator = spotify.NewAuthenticator(
		config.RedirectURL, spotify.ScopePlaylistModifyPublic)
	spotifyAuthenticator.SetAuthInfo(config.ClientID, config.Secret)
	authURL = spotifyAuthenticator.AuthURL(config.State)
	tknHandler = NewTokenHandler()
}

// Client provides an interface to perform actions in Spotify on behalf
// of an authenticated user. Authenticate the user if they haven't already.
func Client(ctx *framework.Context) *SpotifyPlaylister {
	token, found := tknHandler.Get(ctx.Actor.ID)
	if !found {
		content := fmt.Sprintf(StrAuthMessageFmt, authURL)
		ctx.DirectMessage(ctx.Actor.ID, content)

		var err error
		token, err = getToken()
		if err != nil {
			fmt.Printf(
				"Error authorizing Spot for user [%s], %v\n", ctx.Actor, err)
			return nil
		}

		tknHandler.Register(ctx.Actor.ID, token)
	}

	client := spotifyAuthenticator.NewClient(token)
	return &SpotifyPlaylister{client: &client}
}

type SpotifyPlaylister struct {
	client *spotify.Client
}

func (s SpotifyPlaylister) CreatePlaylist(
	userID string,
	playlistName string,
) (*framework.Playlist, error) {
	playlist, err := s.client.CreatePlaylistForUser(
		userID, playlistName, "", true)
	if err != nil {
		return nil, err
	}

	return &framework.Playlist{
		ID:  playlist.ID.String(),
		URL: playlist.ExternalURLs["spotify"],
	}, nil
}

func (s SpotifyPlaylister) CurrentUserID() (*string, error) {
	user, err := s.client.CurrentUser()
	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}
