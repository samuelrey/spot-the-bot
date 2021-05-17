package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	dg.AddHandler(handleUserOpt)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

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

	fmt.Println("Spot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func messageStartUser(s *discordgo.Session, users *[]*discordgo.User, channelID string) (*discordgo.Message, error) {
	u := (*users)[0]
	*users = append((*users)[1:], u)
	m := u.Mention()
	msg := fmt.Sprintf("%v, it's your turn to start the playlist!", m)
	return s.ChannelMessageSend(channelID, msg)
}

func handleUserOpt(s *discordgo.Session, m *discordgo.MessageCreate) {
	// how to check whether message is in specific channel.
	// TODO: replace this with context
	if m.ChannelID != "" {
		fmt.Println("ignore!")
		return
	}

	if m.Content == "optin" {
		fmt.Println("new user: ", m.Author.ID)
	} else if m.Content == "optout" {
		fmt.Println("remove user: ", m.Author.ID)
	}
}
