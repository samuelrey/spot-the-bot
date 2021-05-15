package main

import (
	"fmt"

	"./framework"

	"github.com/bwmarrin/discordgo"
)

var (
	config  *framework.Config
	userIDs []string
)

func init() {
	config = framework.LoadConfig("secrets.json")
	userIDs = framework.LoadUsers("users.json")
}

func main() {
	dg, err := discordgo.New("Bot " + config.Token)
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
	for _, id := range userIDs {
		u, err := dg.User(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		users = append(users, u)
	}

	_, err = messageStartUser(dg, &users, config.ChannelID)
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
