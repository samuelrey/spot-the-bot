package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/samuelrey/spot-discord/cmd"
	"github.com/samuelrey/spot-discord/discord"
	"github.com/samuelrey/spot-discord/framework"
	"github.com/samuelrey/spot-discord/spotify"

	"github.com/bwmarrin/discordgo"
)

var (
	CmdHandler    *framework.CommandHandler
	discordConfig *discord.Config
	TknHandler    *spotify.TokenHandler
	enrolledUsers = make([]string, 0)
)

const (
	prefix = "!"
)

func init() {
	discordConfig = discord.LoadConfig("secrets_discord.json")
}

func main() {
	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	TknHandler = spotify.NewTokenHandler()

	// Configure discord client.
	discord, err := discordgo.New("Bot " + discordConfig.Token)
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

	// Start server to handle Spotify OAuth callback.
	authServer := spotify.StartAuthServer()

	defer func() {
		if err := authServer.Shutdown(context.Background()); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Authentication server shutdown.")
		}
	}()

	fmt.Println("Spot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println()
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.Bot {
		return
	}

	// TODO remove
	if message.ChannelID != discordConfig.ChannelID {
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
			token, err = spotify.AuthorizeSpotForUser(user.ID)
			if err != nil {
				fmt.Println("Error authorizing Spot, ", err)
				return
			}
			TknHandler.Register(user.ID, token)
		}
		spotifyClient := spotify.SpotifyClient(token)

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

	ctx := framework.NewContext(discord, channel, &enrolledUsers, user)
	c := *command
	c(ctx)
}

func registerCommands() {
	CmdHandler.Register("join", cmd.Join)
	CmdHandler.Register("leave", cmd.Leave)
	CmdHandler.Register("list", cmd.List)
}
