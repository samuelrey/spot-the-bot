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
	ServerID     string `json:"DISCORD_SERVER_ID"`
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

	defer dg.Close()

	members, err := dg.GuildMembers(secret.ServerID, "", 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, m := range members {
		p, err := dg.UserChannelPermissions(m.User.ID, secret.ChannelID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%v: %v\n", m.User.Username, p)
	}
}
