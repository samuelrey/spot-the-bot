package framework

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token               string `json:"DISCORD_TOKEN"`
	ServerID            string `json:"DISCORD_SERVER_ID"`
	ChannelID           string `json:"DISCORD_CHANNEL_ID"`
	SpotifyClientID     string `json:"SPOTIFY_CLIENT_ID"`
	SpotifyClientSecret string `json:"SPOTIFY_CLIENT_SECRET"`
}

func LoadConfig(filename string) *Config {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config, ", err)
		return nil
	}
	var config Config
	json.Unmarshal(body, &config)
	return &config
}
