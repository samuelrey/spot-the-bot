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
func Client(ctx *framework.Context) *spotify.Client {
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
	return &client
}
