package framework

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token        string `json:"DISCORD_TOKEN"`
	ServerID     string `json:"DISCORD_SERVER_ID"`
	ChannelID    string `json:"DISCORD_CHANNEL_ID"`
	PlaylistLink string `json:"SPOTIFY_PLAYLIST"`
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

func LoadUsers(filename string) []string {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading user file, ", err)
		return nil
	}
	var userFile []string
	json.Unmarshal(body, &userFile)
	return userFile
}
