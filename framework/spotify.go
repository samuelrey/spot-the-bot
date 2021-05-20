package framework

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type (
	TknMap       map[string]*oauth2.Token
	TokenHandler struct {
		tnks TknMap
	}
)

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{make(TknMap)}
}

func (tknHandler TokenHandler) Get(key string) (*oauth2.Token, bool) {
	tkn, found := tknHandler.tnks[key]
	return tkn, found
}

func (tknHandler TokenHandler) Register(key string, token *oauth2.Token) {
	tknHandler.tnks[key] = token
}

func NewSpotifyClient(config *Config) *spotify.Client {
	cred := &clientcredentials.Config{
		ClientID:     config.SpotifyClientID,
		ClientSecret: config.SpotifyClientSecret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := cred.Token(context.Background())
	if err != nil {
		fmt.Println("Error producing token, ", err)
		return nil
	}
	client := spotify.Authenticator{}.NewClient(token)
	return &client
}
