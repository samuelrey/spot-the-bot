package spotify

import (
	"fmt"
	"net/http"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var (
	config               *Config
	spotifyAuthenticator spotify.Authenticator
	tokenChan            = make(chan *oauth2.Token)
)

func init() {
	config = LoadConfig("secrets_spotify.json")
	spotifyAuthenticator = spotify.NewAuthenticator(
		config.RedirectURL, spotify.ScopePlaylistModifyPublic)
	spotifyAuthenticator.SetAuthInfo(config.ClientID, config.Secret)
}

func SpotifyClient(token *oauth2.Token) *spotify.Client {
	client := spotifyAuthenticator.NewClient(token)
	return &client
}

func AuthorizeSpotForUser(userID string) (*oauth2.Token, error) {
	// TODO DM the url to the user directly.
	authUrl := spotifyAuthenticator.AuthURL(config.State)
	fmt.Printf("%v, %v\n", userID, authUrl)

	token := <-tokenChan

	if token == nil {
		return nil, fmt.Errorf("didn't get token from Spotify")
	}

	return token, nil
}

// StartAuthServer creates HTTP server to handle callback request from Spotify.
// This function is meant to be called only once.
func StartAuthServer() *http.Server {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/callback", authCallback)

	go func() {
		fmt.Println("Authentication server starting")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return server
}

func authCallback(w http.ResponseWriter, r *http.Request) {
	token, err := spotifyAuthenticator.Token(config.State, r)

	if err != nil {
		select {
		case tokenChan <- nil:
			fmt.Println("Error getting token,", err)
		default:
			// Protect against blocking the goroutine.
			fmt.Println("Error endpoint accessed directly")
		}
		return
	}

	fmt.Println("User authorized Spot.")
	tokenChan <- token
}
