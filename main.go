package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Secret struct {
	Token        string   `json:"DISCORD_TOKEN"`
	ServerID     string   `json:"DISCORD_SERVER_ID"`
	ChannelID    string   `json:"DISCORD_CHANNEL_ID"`
	PlaylistLink string   `json:"SPOTIFY_PLAYLIST"`
	UserIDs      []string `json:"USERS"`
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

	// hardcoded users.
	users := make([]*discordgo.User, 0)
	for _, id := range secret.UserIDs {
		u, err := dg.User(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		users = append(users, u)
	}

	_, err = messageStartUser(dg, &users, secret.ChannelID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func messageStartUser(s *discordgo.Session, users *[]*discordgo.User, channelID string) (*discordgo.Message, error) {
	u := (*users)[0]
	*users = append((*users)[1:], u)
	m := u.Mention()
	msg := fmt.Sprintf("%v, it's your turn to start the playlist!", m)
	return s.ChannelMessageSend(channelID, msg)
}
