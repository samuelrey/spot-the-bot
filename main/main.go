package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"../cmd"
	"../framework"

	"github.com/bwmarrin/discordgo"
)

var (
	CmdHandler    *framework.CommandHandler
	config        *framework.Config
	TknHandler    *framework.TokenHandler
	enrolledUsers = make(map[string]bool)
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

	members, err := discord.GuildMembers(config.ServerID, "", 1000)
	if err != nil {
		fmt.Println("Error retrieving guild members, ", err)
		return
	}
	for _, member := range members {
		if !member.User.Bot {
			enrolledUsers[member.User.ID] = false
		}
	}

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

	// For now only check messages in a specific channel.
	if message.ChannelID != config.ChannelID {
		return
	}

	content := message.Content
	// TODO use a command prefix to filter through messages.
	if len(content) == 0 {
		return
	}

	if content == "auth" {
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

	// TODO split content into command name and arguments.
	command, found := CmdHandler.Get(content)
	if !found {
		return
	}

	channel, err := discord.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error retrieving channel, ", err)
		return
	}

	ctx := framework.NewContext(discord, channel, enrolledUsers, config.PlaylistLink, user)
	c := *command
	c(ctx)
}

func registerCommands() {
	CmdHandler.Register("join", cmd.Join)
	CmdHandler.Register("leave", cmd.Leave)
	CmdHandler.Register("link", cmd.Link)
}
