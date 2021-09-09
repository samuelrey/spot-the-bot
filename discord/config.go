package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DiscordConfig struct {
	Token string `json:"DISCORD_TOKEN"`
}

func LoadConfig(filename string) *DiscordConfig {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config, ", err)
		return nil
	}
	var config DiscordConfig
	json.Unmarshal(body, &config)
	return &config
}

func LoadConfigFromEnv() *DiscordConfig {
	return &DiscordConfig{
		Token: os.Getenv("DISCORD_TOKEN"),
	}
}
