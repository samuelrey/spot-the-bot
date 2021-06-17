package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-discord/framework"
)

// TODO make prefix configurable.
const (
	prefix = "!"
)

type DiscordBuilder struct {
	commandHandler  *framework.CommandHandler
	session         *discordgo.Session
	enrolledUsers   *[]framework.User
	playlistBuilder framework.PlaylistBuilder
}

func NewDiscordBuilder(
	config *Config,
	commandHandler *framework.CommandHandler,
	enrolledUsers *[]framework.User,
	playlistBuilder framework.PlaylistBuilder,
) (*DiscordBuilder, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	d := DiscordBuilder{
		commandHandler:  commandHandler,
		session:         session,
		enrolledUsers:   enrolledUsers,
		playlistBuilder: playlistBuilder,
	}

	d.session.AddHandler(d.handleMessage)
	d.session.Identify.Intents = discordgo.IntentsGuildMessages

	return &d, nil
}

func (d *DiscordBuilder) handleMessage(
	dg *discordgo.Session,
	message *discordgo.MessageCreate,
) {
	user := message.Author
	if user.Bot {
		return
	}

	content := message.Content

	if len(content) <= len(prefix) {
		return
	}

	if content[:len(prefix)] != prefix {
		return
	}

	args := strings.Fields(content[len(prefix):])
	name := strings.ToLower(args[0])

	command, found := d.commandHandler.Get(name)
	if !found {
		return
	}

	ctx := NewContext(dg, message.ChannelID, d.enrolledUsers, user)
	c := *command
	c(ctx)
}
