package spotify

import (
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

type SpotifyBuilder struct {
	authenticator *spotify.Authenticator
	authURL       string
	client        *spotify.Client
	user          *spotify.User
}

func CreateSpotifyBuilder(config *Config) *SpotifyBuilder {
	spotifyAuthenticator := spotify.NewAuthenticator(
		config.RedirectURL,
		spotify.ScopePlaylistModifyPrivate,
	)
	spotifyAuthenticator.SetAuthInfo(config.ClientID, config.Secret)

	authURL = spotifyAuthenticator.AuthURL(config.State)

	return &SpotifyBuilder{
		authenticator: &spotifyAuthenticator,
		authURL:       authURL,
	}
}

// AuthorizeUser takes a user through the Spotify authorization flow. The
// token we receive from Spotify is used to create a reusable client. The
// client is tied to a single Spotify user. We recommend creating an account
// on Spotify specific for this bot.
func (sb *SpotifyBuilder) AuthorizeUser() error {
	log.Printf("Navigate here to authorize Spotify user: %s\n", authURL)

	token, err := getToken()
	if err != nil {
		return errors.Wrap(err, "Authorize Spotify user")
	}

	client := spotifyAuthenticator.NewClient(token)
	spotifyUser, err := client.CurrentUser()
	if err != nil {
		return errors.Wrap(err, "Authorize Spotify user")
	}

	sb.client = &client
	sb.user = &spotifyUser.User
	return nil
}

// CreatePlaylist creates a playlist with the given name for the authorized
// user.
func (sb *SpotifyBuilder) CreatePlaylist(playlistName string) (*framework.Playlist, error) {
	playlist, err := sb.client.CreateCollaborativePlaylistForUser(sb.user.ID, playlistName, "")
	if err != nil {
		return nil, errors.Wrap(err, "Create Spotify playlist")
	}

	return &framework.Playlist{
		ID:  playlist.ID.String(),
		URL: playlist.ExternalURLs["spotify"],
	}, nil
}
