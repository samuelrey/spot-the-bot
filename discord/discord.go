package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-discord/framework"
)

var (
	CmdHandler    *framework.CommandHandler
	config        *Config
	enrolledUsers = make([]framework.User, 0)
)

const (
	prefix = "!"
)

func init() {
	config = LoadConfig("secrets_discord.json")
}

// DiscordSession connects to discord using the bot token. It accepts a
// CommandHandler which informs us how to handle different commands.
func DiscordSession(cmdHandler *framework.CommandHandler) (*discordgo.Session, error) {
	CmdHandler = cmdHandler

	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	err = session.Open()
	if err != nil {
		return nil, err
	}

	session.AddHandler(handleMessage)
	session.Identify.Intents = discordgo.IntentsGuildMessages
	return session, nil
}

func handleMessage(dg *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.Bot {
		return
	}

	// TODO remove to enable spot on all channels.
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

	command, found := CmdHandler.Get(name)
	if !found {
		return
	}

	channel, err := dg.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error retrieving channel, ", err)
		return
	}

	ctx := NewContext(dg, channel, &enrolledUsers, user)
	c := *command
	c(ctx)
}
