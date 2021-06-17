package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-discord/framework"
)

type DiscordBuilder struct {
	commandHandler *framework.CommandHandler
	session        *discordgo.Session
}

func CreateDiscordBuilder(
	config *Config,
	commandHandler *framework.CommandHandler,
) (*DiscordBuilder, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	return &DiscordBuilder{
		commandHandler: commandHandler,
		session:        session,
	}, nil
}

func (d *DiscordBuilder) Reply(channelID, content string) error {
	_, err := d.session.ChannelMessageSend(channelID, content)
	return err
}

func (d *DiscordBuilder) DirectMessage(recipientID, content string) error {
	userChannel, err := d.session.UserChannelCreate(recipientID)
	if err != nil {
		return err
	}

	_, err = d.session.ChannelMessageSend(userChannel.ID, content)
	return err
}

func (d *DiscordBuilder) Open() error {
	d.session.AddHandler(d.handleMessage)
	d.session.Identify.Intents = discordgo.IntentsGuildMessages

	return d.session.Open()
}

func (d *DiscordBuilder) Close() error {
	return d.session.Close()
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

	command, found := CmdHandler.Get(name)
	if !found {
		return
	}

	ctx := NewContext(dg, message.ChannelID, &enrolledUsers, user)
	c := *command
	c(ctx)
}