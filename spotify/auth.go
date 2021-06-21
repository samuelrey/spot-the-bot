package spotify

import (
	"context"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/samuelrey/spot-discord/framework"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var tokenChan = make(chan *oauth2.Token)
var errChan = make(chan error)

type SpotifyAuthorizer struct {
	authenticator *spotify.Authenticator
	authURL       string
	state         string
}

func NewSpotifyAuthorizer(config *Config) *SpotifyAuthorizer {
	spotifyAuthenticator := spotify.NewAuthenticator(
		config.RedirectURL,
		spotify.ScopePlaylistModifyPrivate,
	)
	spotifyAuthenticator.SetAuthInfo(config.ClientID, config.Secret)

	authURL := spotifyAuthenticator.AuthURL(config.State)

	return &SpotifyAuthorizer{
		authenticator: &spotifyAuthenticator,
		authURL:       authURL,
		state:         config.State,
	}
}

// AuthorizeUser takes a user through the Spotify authorization flow. Return
// a PlaylistCreator which is tied to a single Spotify user's account. We
// recommend creating an account on Spotify specific for this bot.
func (sa *SpotifyAuthorizer) AuthorizeUser() (framework.PlaylistCreator, error) {
	log.Println("Auth server starting.")
	srv := sa.startAuthServer()

	defer func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println(err)
		} else {
			log.Println("Auth server shutdown.")
		}
	}()

	log.Printf("Navigate here to authorize Spotify user: %s\n", sa.authURL)

	token, err := getToken()
	if err != nil {
		return nil, errors.Wrap(err, "authorize spotify user")
	}

	client := sa.authenticator.NewClient(token)
	spotifyUser, err := client.CurrentUser()
	if err != nil {
		return nil, errors.Wrap(err, "authorize spotify user")
	}

	return &SpotifyPlaylistCreator{
		client: &client,
		user:   &spotifyUser.User,
	}, nil
}

// startAuthServer creates HTTP server to handle callback request from Spotify.
func (sa SpotifyAuthorizer) startAuthServer() *http.Server {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/callback", sa.authCallback)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return server
}

// authCallback writes the Spotify token or error to their respective channels.
func (sa SpotifyAuthorizer) authCallback(w http.ResponseWriter, r *http.Request) {
	token, err := sa.authenticator.Token(sa.state, r)

	if err != nil {
		errChan <- err
		return
	}

	tokenChan <- token
}

// getToken reads the Spotify token or error written by
func getToken() (*oauth2.Token, error) {
	select {
	case token := <-tokenChan:
		return token, nil
	case err := <-errChan:
		return nil, err
	}
}
