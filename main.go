package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	go dummy()

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

func dummy() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			u := userIDs[0]
			userIDs = append(userIDs[1:], u)
			fmt.Printf("%v, it's your turn!", u)
		}
	}
}

func handleUserOpt(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != config.ChannelID {
		fmt.Println("ignore!")
		return
	}

	user := m.Author
	if m.Content == "optin" {
		enrolled := false
		for _, id := range userIDs {
			if id == user.ID {
				enrolled = true
				break
			}
		}
		if !enrolled {
			userIDs = append(userIDs, user.ID)
			fmt.Println("new user: ", user.ID)
		}
	} else if m.Content == "optout" {
		found := -1
		for i, id := range userIDs {
			if id == user.ID {
				found = i
				break
			}
		}
		if found != -1 {
			userIDs = append(userIDs[:found], userIDs[found+1:]...)
			fmt.Println("remove user: ", m.Author.ID)
		}
	}
	fmt.Println(userIDs)
}
