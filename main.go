package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Secret struct {
	Token        string `json:"DISCORD_TOKEN"`
	ChannelID    string `json:"DISCORD_CHANNEL_ID"`
	PlaylistLink string `json:"SPOTIFY_PLAYLIST"`
}

func main() {
	f, err := os.Open("secrets.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(f)
	var secret Secret
	json.Unmarshal(byteValue, &secret)

	dg, err := discordgo.New("Bot " + secret.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	dg.ChannelMessageSend(secret.ChannelID, secret.PlaylistLink)
	dg.Close()
}
