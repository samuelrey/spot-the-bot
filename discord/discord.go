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
	DiscordConnector
	DiscordMessager
	commandHandler  *framework.CommandHandler
	enrolledUsers   *[]framework.User
	playlistBuilder framework.PlaylistCreator
}

func NewDiscordBuilder(
	config *Config,
	commandHandler *framework.CommandHandler,
	enrolledUsers *[]framework.User,
	playlistBuilder framework.PlaylistCreator,
) (*DiscordBuilder, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	d := DiscordBuilder{
		DiscordConnector: DiscordConnector{
			session: session,
		},
		DiscordMessager: DiscordMessager{
			session: session,
		},
		commandHandler:  commandHandler,
		enrolledUsers:   enrolledUsers,
		playlistBuilder: playlistBuilder,
	}

	session.AddHandler(d.handleMessage)
	session.Identify.Intents = discordgo.IntentsGuildMessages

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

type DiscordMessager struct {
	session *discordgo.Session
}

func (dm *DiscordMessager) Reply(channelID, content string) error {
	_, err := dm.session.ChannelMessageSend(channelID, content)
	return err
}

func (dm *DiscordMessager) DirectMessage(recipientID, content string) error {
	userChannel, err := dm.session.UserChannelCreate(recipientID)
	if err != nil {
		return err
	}

	_, err = dm.session.ChannelMessageSend(userChannel.ID, content)
	return err
}

type DiscordConnector struct {
	session *discordgo.Session
}

func (dc *DiscordConnector) Open() error {
	return dc.session.Open()
}

func (dc *DiscordConnector) Close() error {
	return dc.session.Close()
}
