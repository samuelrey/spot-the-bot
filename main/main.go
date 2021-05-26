package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/samuelrey/spot-discord/cmd"
	"github.com/samuelrey/spot-discord/framework"

	"github.com/bwmarrin/discordgo"
)

var (
	CmdHandler    *framework.CommandHandler
	config        *framework.Config
	TknHandler    *framework.TokenHandler
	enrolledUsers = make([]string, 0)
)

const (
	prefix = "!"
)

func init() {
	config = framework.LoadConfig("secrets.json")
}

func main() {
	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	TknHandler = framework.NewTokenHandler()

	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer discord.Close()

	discord.AddHandler(commandHandler)
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	fmt.Println("Spot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.Bot {
		return
	}

	// TODO remove
	if message.ChannelID != config.ChannelID {
		return
	}

	content := message.Content

	// Check message for command prefix.
	if len(content) <= len(prefix) {
		return
	}

	if content[:len(prefix)] != prefix {
		return
	}

	args := strings.Fields(content[len(prefix):])
	name := strings.ToLower(args[0])

	// TODO remove
	if name == "auth" {
		token, found := TknHandler.Get(user.ID)
		if !found {
			var err error
			token, err = framework.AuthorizeSpotForUser(user.ID)
			if err != nil {
				fmt.Println("Error authorizing Spot, ", err)
				return
			}
			TknHandler.Register(user.ID, token)
		}
		spotifyClient := framework.SpotifyClient(token)

		// Verify we got a good token.
		u, err := spotifyClient.CurrentUser()
		if err != nil {
			fmt.Println("Error using spotify client, ", err)
			return
		}
		fmt.Println(u.ID)
	}

	command, found := CmdHandler.Get(name)
	if !found {
		return
	}

	channel, err := discord.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error retrieving channel, ", err)
		return
	}

	ctx := framework.NewContext(
		discord, channel, &enrolledUsers, config.PlaylistLink, user)
	c := *command
	c(ctx)
}

func registerCommands() {
	CmdHandler.Register("join", cmd.Join)
	CmdHandler.Register("leave", cmd.Leave)
	CmdHandler.Register("link", cmd.Link)
	CmdHandler.Register("list", cmd.List)
}
