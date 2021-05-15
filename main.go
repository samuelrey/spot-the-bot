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

	// Sending a message
	// dg.ChannelMessageSend(secret.ChannelID, secret.PlaylistLink)

	// Trying to understand channel-user relationships
	// channels, _ := dg.GuildChannels(secret.ServerID)
	// for _, c := range channels {
	// 	fmt.Printf("%+v\n", c)
	// 	for _, p := range c.PermissionOverwrites {
	// 		fmt.Printf("%+v\n", p)
	// 	}
	// }

	// fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-sc
}
