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
	session         *discordgo.Session
	commandHandler  *framework.CommandHandler
	enrolledUsers   *[]framework.MessageUser
	playlistBuilder framework.PlaylistCreator
}

func NewDiscordBuilder(
	config *Config,
	commandHandler *framework.CommandHandler,
	enrolledUsers *[]framework.MessageUser,
	playlistBuilder framework.PlaylistCreator,
) (*DiscordBuilder, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	d := DiscordBuilder{
		session:         session,
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

	// TODO get playlist name from config
	ctx := framework.CommandContext{
		Messager: &DiscordMessager{
			session:   dg,
			channelID: message.ChannelID,
		},
		PlaylistCreator: d.playlistBuilder,
		PlaylistName:    "Einstok",
		EnrolledUsers:   d.enrolledUsers,
		Actor: framework.MessageUser{
			ID:       user.ID,
			Username: user.Username,
		},
	}
	c := *command
	c(&ctx)
}

func (db *DiscordBuilder) Open() error {
	return db.session.Open()
}

func (db *DiscordBuilder) Close() error {
	return db.session.Close()
}

type DiscordMessager struct {
	session   *discordgo.Session
	channelID string
}

func (dm *DiscordMessager) Reply(content string) error {
	_, err := dm.session.ChannelMessageSend(dm.channelID, content)
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
