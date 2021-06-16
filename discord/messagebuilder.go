package discord

import "github.com/bwmarrin/discordgo"

type DiscordBuilder struct {
	session *discordgo.Session
}

func CreateDiscordBuilder(config *Config) (*DiscordBuilder, error) {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	return &DiscordBuilder{session: session}, nil
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