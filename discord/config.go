package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token string `json:"DISCORD_TOKEN"`
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
