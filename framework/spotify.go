package framework

import (
	"context"
	"fmt"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

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
