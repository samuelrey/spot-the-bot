package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Messager struct {
	Session   *discordgo.Session
	ChannelID string
}

// Reply sends a message with the given contents to the channel where the
// command was received.
func (dm *Messager) Reply(content string) error {
	_, err := dm.Session.ChannelMessageSend(dm.ChannelID, content)
	return err
}

// DirectMessage sends a message with the given contents to a Discord user in
// private.
func (dm *Messager) DirectMessage(recipientID, content string) error {
	userChannel, err := dm.Session.UserChannelCreate(recipientID)
	if err != nil {
		return err
	}

	_, err = dm.Session.ChannelMessageSend(userChannel.ID, content)
	return err
}
