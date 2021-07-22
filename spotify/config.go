package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type SpotifyConfig struct {
	ClientID    string `json:"CLIENT_ID"`
	RedirectURL string `json:"REDIRECT_URL"`
	Secret      string `json:"SECRET"`
}

func LoadConfig(filename string) *SpotifyConfig {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config, ", err)
		return nil
	}
	var config SpotifyConfig
	json.Unmarshal(body, &config)
	return &config
}

func LoadConfigFromEnv() SpotifyConfig {
	return SpotifyConfig{
		ClientID:    os.Getenv("SPOTIFY_CLIENT_ID"),
		RedirectURL: os.Getenv("SPOTIFY_REDIRECT_URL"),
		Secret:      os.Getenv("SPOTIFY_SECRET"),
	}
}
