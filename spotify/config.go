package spotify

import (
	"os"
)

type SpotifyConfig struct {
	ClientID    string
	RedirectURL string
	Secret      string
}

func LoadConfig() SpotifyConfig {
	return SpotifyConfig{
		ClientID:    os.Getenv("SPOTIFY_CLIENT_ID"),
		RedirectURL: os.Getenv("SPOTIFY_REDIRECT_URL"),
		Secret:      os.Getenv("SPOTIFY_SECRET"),
	}
}
