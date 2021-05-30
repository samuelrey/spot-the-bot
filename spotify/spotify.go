package spotify

import (
	"fmt"

	"github.com/zmb3/spotify"
)

var (
	authURL              string
	config               *Config
	spotifyAuthenticator spotify.Authenticator
	tknHandler           *TokenHandler
)

func init() {
	config = LoadConfig("secrets_spotify.json")
	spotifyAuthenticator = spotify.NewAuthenticator(
		config.RedirectURL, spotify.ScopePlaylistModifyPublic)
	spotifyAuthenticator.SetAuthInfo(config.ClientID, config.Secret)
	authURL = spotifyAuthenticator.AuthURL(config.State)
	tknHandler = NewTokenHandler()
}

// SpotifyClient provides an interface to perform actions in Spotify on behalf
// of an authenticated user.
func SpotifyClient(userID string) *spotify.Client {
	token, found := tknHandler.Get(userID)
	if !found {
		// TODO DM the url to the user directly.
		fmt.Printf("%v, %v\n", userID, authURL)

		var err error
		token, err = getToken()
		if err != nil {
			fmt.Println("Error authorizing Spot, ", err)
			return nil
		}

		tknHandler.Register(userID, token)
	}

	client := spotifyAuthenticator.NewClient(token)
	return &client
}
