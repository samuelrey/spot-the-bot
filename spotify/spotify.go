package spotify

import (
	"context"
	"log"

	"github.com/pkg/errors"
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

type SpotifyAuthorizer struct {
	authenticator *spotify.Authenticator
	authURL       string
}

type SpotifyPlaylistCreator struct {
	client *spotify.Client
	user   *spotify.User
}

func NewSpotifyAuthorizer(config *Config) *SpotifyAuthorizer {
	spotifyAuthenticator := spotify.NewAuthenticator(
		config.RedirectURL,
		spotify.ScopePlaylistModifyPrivate,
	)
	spotifyAuthenticator.SetAuthInfo(config.ClientID, config.Secret)

	authURL = spotifyAuthenticator.AuthURL(config.State)

	return &SpotifyAuthorizer{
		authenticator: &spotifyAuthenticator,
		authURL:       authURL,
	}
}

// AuthorizeUser takes a user through the Spotify authorization flow. The
// token we receive from Spotify is used to create a reusable client. The
// client is tied to a single Spotify user. We recommend creating an account
// on Spotify specific for this bot.
func (sa *SpotifyAuthorizer) AuthorizeUser() (framework.PlaylistCreator, error) {
	log.Println("Authentication server starting")
	srv := StartAuthServer()

	defer func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println(err)
		} else {
			log.Println("Authentication server shutdown.")
		}
	}()

	log.Printf("Navigate here to authorize Spotify user: %s\n", authURL)

	token, err := getToken()
	if err != nil {
		return nil, errors.Wrap(err, "Authorize Spotify user")
	}

	client := spotifyAuthenticator.NewClient(token)
	spotifyUser, err := client.CurrentUser()
	if err != nil {
		return nil, errors.Wrap(err, "Authorize Spotify user")
	}

	return &SpotifyPlaylistCreator{
		client: &client,
		user:   &spotifyUser.User,
	}, nil
}

// CreatePlaylist creates a playlist with the given name for the authorized
// user.
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
