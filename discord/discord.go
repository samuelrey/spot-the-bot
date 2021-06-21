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

// handleMessage reads messages sent in the guild and runs commands based on
// those messages.
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
	// TODO optionally populate dependencies based on command
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

// Open is a wrapper function to open a Discord session.
func (db *DiscordBuilder) Open() error {
	return db.session.Open()
}

// Close is a wrapper function to close a Discord session.
func (db *DiscordBuilder) Close() error {
	return db.session.Close()
}

type DiscordMessager struct {
	session   *discordgo.Session
	channelID string
}

// Reply sends a message with the given contents to the channel where the
// command was received.
func (dm *DiscordMessager) Reply(content string) error {
	_, err := dm.session.ChannelMessageSend(dm.channelID, content)
	return err
}

// DirectMessage sends a message with the given contents to a Discord user in
// private.
func (dm *DiscordMessager) DirectMessage(recipientID, content string) error {
	userChannel, err := dm.session.UserChannelCreate(recipientID)
	if err != nil {
		return err
	}

	_, err = dm.session.ChannelMessageSend(userChannel.ID, content)
	return err
}
