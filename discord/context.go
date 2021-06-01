package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-discord/framework"
)

// NewContext takes discord specific data and produces a Context which can be
// used by commands.
func NewContext(
	session *discordgo.Session,
	channel *discordgo.Channel,
	enrolledUsers *[]framework.User,
	actor *discordgo.User,
) *framework.Context {
	ctx := new(framework.Context)
	discordMessager := DiscordMessager{
		Session: session,
		Channel: channel,
	}
	ctx.Messager = discordMessager
	ctx.EnrolledUsers = enrolledUsers
	ctx.Actor = framework.User{
		ID:       actor.ID,
		Username: actor.Username,
	}
	return ctx
}

type DiscordMessager struct {
	Session *discordgo.Session
	Channel *discordgo.Channel
}

func (d DiscordMessager) Reply(content string) error {
	_, err := d.Session.ChannelMessageSend(d.Channel.ID, content)
	return err
}

func (d DiscordMessager) DirectMessage(recipient, content string) error {
	channel, err := d.Session.UserChannelCreate(recipient)
	if err != nil {
		return err
	}

	_, err = d.Session.ChannelMessageSend(channel.ID, content)
	return err
}
