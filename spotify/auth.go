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

// TODO create error channel.
var tokenChan = make(chan *oauth2.Token)

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

// AuthorizeUser takes a user through the Spotify authorization flow. The
// token we receive from Spotify is used to create a reusable client. The
// client is tied to a single Spotify user. We recommend creating an account
// on Spotify specific for this bot.
func (sa *SpotifyAuthorizer) AuthorizeUser() (framework.PlaylistCreator, error) {
	log.Println("Authentication server starting")
	srv := sa.startAuthServer()

	defer func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println(err)
		} else {
			log.Println("Authentication server shutdown.")
		}
	}()

	log.Printf("Navigate here to authorize Spotify user: %s\n", sa.authURL)

	token, err := getToken()
	if err != nil {
		return nil, errors.Wrap(err, "Authorize Spotify user")
	}

	client := sa.authenticator.NewClient(token)
	spotifyUser, err := client.CurrentUser()
	if err != nil {
		return nil, errors.Wrap(err, "Authorize Spotify user")
	}

	return &SpotifyPlaylistCreator{
		client: &client,
		user:   &spotifyUser.User,
	}, nil
}

// startAuthServer creates HTTP server to handle callback request from Spotify.
// This function is meant to be called only once.
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

func (sa SpotifyAuthorizer) authCallback(w http.ResponseWriter, r *http.Request) {
	token, err := sa.authenticator.Token(sa.state, r)

	if err != nil {
		select {
		case tokenChan <- nil:
			log.Println("Error getting token:", err)
		default:
			// Protect against blocking the goroutine.
			log.Println("Error endpoint accessed directly")
		}
		return
	}

	select {
	case tokenChan <- token:
		log.Println("User authorized Spot")
	default:
		// Protect against blocking the goroutine.
		log.Println("User already authorized Spot")
	}
}

func getToken() (*oauth2.Token, error) {
	token := <-tokenChan
	if token == nil {
		return nil, errors.New("token not received")
	}

	return token, nil
}
